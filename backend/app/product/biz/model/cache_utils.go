package model

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"
)

const (
	cacheDuration      = 1 * time.Hour    // 正常缓存时间
	emptyCacheDuration = 30 * time.Minute // 空值缓存时间（防穿透）
)

var (
	cacheEmptyErr = errors.New("cache empty value") // 缓存为空的错误
)

var (
	ProductLocks     *sync.Map // 数据库全局锁的实例，缓存击穿时，只放行一个进入 MYSQL
	onceProductLocks sync.Once
)

func GetProductLocks() *sync.Map {
	onceProductLocks.Do(func() {
		ProductLocks = &sync.Map{}
	})
	return ProductLocks
}

func getProductListCacheKey(query, category string) string {
	key := "product_list:"
	if query != "" {
		key += "query=" + query
	}
	if category != "" {
		key += "_category=" + category
	}
	return key
}

func getProductDetailCacheKey(productId int64) string {
	return fmt.Sprintf("product_detail:%d", productId)
}

func getCachedExpiration(baseDuration, randomDuration time.Duration) time.Duration {
	// 生成一个随机的时间偏移量
	offset := time.Duration(randomDuration) * time.Second
	// 返回基础时间加上随机偏移量的结果
	return baseDuration + offset
}

// 缓存查询结果。
// key 是缓存的键
// value 是一个指针类型，用于接收反序列化后的结果
// 若返回的 err = cacheEmptyErr，表示缓存空值
func getValueByKeyFromCache(ctx context.Context, redisCli *redis.Client, key string, value any) (err error) {
	cachedResult := redisCli.Get(ctx, key)

	err = cachedResult.Err()
	if err != nil {
		return err
	}

	// 获取字节流
	cachedBytes, err := cachedResult.Bytes()
	if err != nil {
		return err
	}

	if len(cachedBytes) == 0 {
		// 空值表示缓存穿透保护
		return cacheEmptyErr
	}
	err = json.Unmarshal(cachedBytes, &value)
	if err != nil {
		return err
	}
	return nil
}
