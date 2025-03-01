package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"github.com/tiktokmall/backend/app/user/infra/mq"
	"gorm.io/gorm"
)

type UserDAO struct {
	ctx   context.Context
	db    *gorm.DB
	redis *redis.Client
	locks *sync.Map // 用于缓存击穿防护的互斥锁
	// bloomFilter *BloomFilter // 布隆过滤器
}

func NewUserDAO(ctx context.Context, db *gorm.DB, redis *redis.Client) *UserDAO {
	return &UserDAO{
		ctx:   ctx,
		db:    db,
		redis: redis,
		locks: GetUserLocks(),
	}
}

func (dao *UserDAO) Create(user *User) (err error) {
	// TODO: 请实现Create方法
	// 插入数据库
	err = dao.db.WithContext(dao.ctx).Model(&User{}).Create(&user).Error
	if err != nil {
		return err
	}
	// 插入缓存
	cachedKey := getUserEmailKey(user.Email)
	randomExpiration := getCachedExpiration(cacheDuration, emptyCacheDuration)
	userData, err := json.Marshal(user)
	if err != nil {
		hlog.Errorf("设置缓存，marshal user data failed: %v", err)
		return nil
	}
	dao.redis.Set(dao.ctx, cachedKey, userData, randomExpiration)
	return nil
}

func (dao *UserDAO) GetByEmail(email string) (*User, error) {
	// TODO: 请实现GetByEmail方法
	// 从缓存中获取
	cachedKey := getUserEmailKey(email)
	// 第一次缓存检查
	val, err := dao.redis.Get(dao.ctx, cachedKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("user not found (cached)")
		}
		var cachedUser User
		if err := json.Unmarshal([]byte(val), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}

	// 缓存未找到，获取锁，查询数据库。防止缓存击穿
	// 根据 cachedKey 获取锁
	mu, _ := dao.locks.LoadOrStore(cachedKey, &sync.Mutex{})
	lock := mu.(*sync.Mutex)

	// 使用TryLock实现带超时的锁获取
	start := time.Now()
	// TODO: 可以考虑时用 channel，需要注释协程是否正确关闭
	for {
		if lock.TryLock() {
			defer lock.Unlock() // 确保最终解锁
			break
		}
		if time.Since(start) > 10*time.Second {
			return nil, fmt.Errorf("系统繁忙，请稍后重试")
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 可以访问数据库。
	// 第二次缓存检查（获得锁后再次检查）
	if val, err = dao.redis.Get(dao.ctx, cachedKey).Result(); err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("user not found (cached)")
		}
		var cachedUser User
		if err := json.Unmarshal([]byte(val), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}
	// 查询数据库
	var user User
	err = dao.db.WithContext(dao.ctx).Where("email = ? and deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 未注册的用户，缓存空值防止缓存穿透，随机过期时间防雪崩
			dao.redis.Set(dao.ctx, cachedKey, "", getCachedExpiration(emptyCacheDuration, emptyCacheDuration))
		}
		return nil, err
	}

	// 写入缓存（带随机过期时间防雪崩）
	if userData, err := json.Marshal(user); err == nil {
		randomExpiration := getCachedExpiration(cacheDuration, emptyCacheDuration)
		dao.redis.Set(context.Background(), cachedKey, userData, randomExpiration)
	}

	return &user, nil
}

// Update 更新用户信息，目前只有用户名和头像可以更新
// 一致性策略：先更新数据库，再异步删除缓存
// 1. 先更新数据库
// 2. 成功后，将删除缓存的操作加入消息队列
// 3. 通过消息队列的重试机制确保缓存最终被删除
func (dao *UserDAO) UpdateUsernameOrAvatorById(user User) (err error) {
	// TODO：可以先根据 email 查询用户，判断是否需要更新。
	// 更新数据库
	err = dao.db.WithContext(dao.ctx).Model(&User{}).Where("id = ? and deleted_at IS NULL", user.ID).Updates(map[string]interface{}{
		"username": user.Username,
		"avator":   user.Avator,
	}).Error
	if err != nil {
		return err
	}

	// 将删除缓存的操作加入消息队列, 异步删除缓存
	cachedKey := getUserEmailKey(user.Email)
	if err = mq.Natscp.AddDeleteCacheMessage(cachedKey); err != nil {
		log.Printf("nats: Failed to add delete cache message, key:%v, err %v", cachedKey, err)
	}
	return nil
}

func (dao *UserDAO) DeleteByID(id int) (err error) {
	// TODO: 请实现DeleteByID方法
	// 查缓存，防止缓存穿透
	cacheDeleteKey := "delete_user_by_id:" + fmt.Sprintf("%d", id)
	if val, err := dao.redis.Get(dao.ctx, cacheDeleteKey).Result(); err == nil {
		if val == "" { // 空值表示已删除，缓存穿透保护
			return nil
		}
	}
	var user User
	err = dao.db.WithContext(dao.ctx).Where("id = ? and deleted_at IS NULL", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 缓存空值防止穿透，随机过期时间防雪崩
			dao.redis.Set(dao.ctx, cacheDeleteKey, "", getCachedExpiration(emptyCacheDuration, emptyCacheDuration))
			return nil
		}
		return err
	}
	err = dao.db.WithContext(dao.ctx).Model(&User{}).Where("id = ?", user.ID).Set("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	// 将删除缓存的操作加入消息队列, 异步删除缓存
	userEmailCachedKey := getUserEmailKey(user.Email)
	if err = mq.Natscp.AddDeleteCacheMessage(userEmailCachedKey); err != nil {
		log.Printf("nats: Failed to add delete cache message, key:%v, err %v", userEmailCachedKey, err)
	}
	return nil
}
