package model

import (
	"sync"
	"time"
)

const (
	cacheDuration      = 1 * time.Hour    // 正常缓存时间
	emptyCacheDuration = 30 * time.Minute // 空值缓存时间（防穿透）
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

func getCachedExpiration(baseDuration, randomDuration time.Duration) time.Duration {
	// 生成一个随机的时间偏移量
	offset := time.Duration(randomDuration) * time.Second
	// 返回基础时间加上随机偏移量的结果
	return baseDuration + offset
}
