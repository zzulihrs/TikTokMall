package model

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/tiktokmall/backend/app/merchant/infra/mq"
)

// 目前实现的 product 相关 key
// - product_detail:{pid}
// - product_list:{query_condition}

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

func ClearProductListCachedKey(ctx context.Context, redisCli *redis.Client) error {
	prefix := "product_list:"
	// 使用 SCAN 命令迭代查找匹配的键
	// var keys []string
	var cursor uint64
	for {
		var err error
		var result []string
		result, cursor, err = redisCli.Scan(ctx, cursor, prefix+"*", 0).Result()
		if err != nil {
			return fmt.Errorf("扫描 Redis 键时出错: %v", err)
		}
		for _, key := range result {
			mq.Natscp.AddDeleteCacheMessage(key)
		}
		if cursor == 0 {
			break
		}
	}

	return nil
}
