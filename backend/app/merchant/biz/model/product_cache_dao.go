package model

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 暂时没用
type ProductDAO struct {
	ctx   context.Context
	db    *gorm.DB
	redis *redis.Client
	locks *sync.Map // 用于缓存击穿防护的互斥锁
	// bloomFilter *BloomFilter // 布隆过滤器
}

func NewProductDAO(ctx context.Context, db *gorm.DB, redis *redis.Client) *ProductDAO {
	return &ProductDAO{
		ctx:   ctx,
		db:    db,
		redis: redis,
		locks: GetProductLocks(),
	}
}
