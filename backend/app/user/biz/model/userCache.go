package model

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	cacheDuration      = 30 * time.Minute // 正常缓存时间
	emptyCacheDuration = 5 * time.Minute  // 空值缓存时间（防穿透）
	cacheKeyPrefix     = "user_email:"    // Redis键前缀
)

type CacheddUserQuery struct {
	ctx   context.Context
	db    *gorm.DB
	redis *redis.Client
	nc    *nats.Conn
	mu    sync.Map // 用于存储每个email对应的互斥锁
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, redis *redis.Client, nc *nats.Conn) *CacheddUserQuery {
	return &CacheddUserQuery{
		ctx:   ctx,
		db:    db,
		redis: redis,
		nc:    nc,
	}
}

func (c *CacheddUserQuery) GetByEmail(ctx context.Context, email string) (*User, error) {
	cacheKey := cacheKeyPrefix + email

	// 第一次缓存检查
	val, err := c.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("user not found (cached)")
		}
		var cachedUser User
		if err := json.Unmarshal([]byte(val), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}

	// 带超时的锁获取
	mu, _ := c.mu.LoadOrStore(email, &sync.Mutex{})
	lock := mu.(*sync.Mutex)

	lockCh := make(chan struct{})
	go func() {
		lock.Lock()
		close(lockCh) // 获取到锁
	}()

	select {
	case <-lockCh:
	case <-time.After(10 * time.Second):
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}
	defer lock.Unlock()

	// 第二次缓存检查（获得锁后再次检查）
	if val, err = c.redis.Get(ctx, cacheKey).Result(); err == nil {
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
	err = c.db.WithContext(ctx).Where("email = ? and deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 缓存空值防止穿透
			c.redis.Set(ctx, cacheKey, "", emptyCacheDuration)
		}
		return nil, err
	}

	// 更新缓存
	userData, _ := json.Marshal(user)
	c.redis.Set(ctx, cacheKey, string(userData), cacheDuration)
	return &user, nil
}

// 新增消息发布结构体
type cacheDeleteMsg struct {
	Email string `json:"email"`
	Delay int    `json:"delay"` // 单位：秒
}

func (c *CacheddUserQuery) UpdateUser(ctx context.Context, user *User) error {
	// 更新数据库（保持原逻辑）

	err := c.db.WithContext(ctx).Model(user).Where("deleted_at IS NULL").Updates(map[string]interface{}{
		"username": user.Username,
		"avator":   user.Avator,
	}).Error
	if err != nil {
		return err
	}
	// 延迟双删逻辑（新增部分）
	cacheKey := cacheKeyPrefix + user.Email

	// 第一次立即删除
	c.redis.Del(ctx, cacheKey)

	// 通过NATS发送延迟删除消息
	msg := cacheDeleteMsg{
		Email: user.Email,
		Delay: 1, // 1秒后二次删除
	}
	msgData, _ := json.Marshal(msg)

	// 使用JetStream发送延迟消息（需要项目已配置JetStream）
	js, _ := c.nc.JetStream()
	js.Publish("USER_CACHE.DELETE", msgData, nats.MsgId(uuid.NewString()))

	return nil
}
