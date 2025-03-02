package mq

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

// CacheMessage 缓存操作消息
type CacheMessage struct {
	Key       string    // 需要删除的缓存key
	Operation string    // 操作类型
	Attempts  int       // 重试次数
	CreatedAt time.Time // 创建时间
}

const (
	cacheProductSubject       = "cache.product"
	maxRetryAttempts          = 3               // 最大重试次数
	retryDelay                = time.Second * 2 // 重试间隔
	messageQueueDeleteListKey = "delete_list"   // 消息队列的中删除 Operation
)

type NatsCacheProcessor struct {
	nc    *nats.Conn
	redis *redis.Client
}

func InitNatsCacheProcessor(nc *nats.Conn, redis *redis.Client) *NatsCacheProcessor {
	return &NatsCacheProcessor{
		nc:    nc,
		redis: redis,
	}
}

// StartProcessor 启动 NATS 消息处理器
// 1. 订阅 cache.product 主题
// 2. 接收到消息后进行缓存删除操作
// 3. 如果删除失败，在当前 goroutine 中进行重试
// 4. 达到最大重试次数后，记录错误日志
func (p *NatsCacheProcessor) StartProcessor() {
	// 订阅 cache.user 主题
	// p.nc.Subscribe(cacheProductSubject, func(m *nats.Msg) {
	// 	// 解析消息
	// 	var msg CacheMessage
	// 	if err := json.Unmarshal(m.Data, &msg); err != nil {
	// 		log.Printf("nats: Failed to unmarshal message: %v", err)
	// 		return
	// 	}
	// 	if msg.Operation == messageQueueDeleteListKey { // 目前 cache.product 主题，只有一个 delete 操作
	// 		// 在当前 goroutine 中进行重试
	// 		// 不需要重新发布消息，因为这会导致消息重复和处理混乱
	// 		var lastErr error
	// 		for attempt := 0; attempt < maxRetryAttempts; attempt++ {
	// 			// 尝试删除缓存
	// 			if err := p.redis.Del(context.Background(), msg.Key).Err(); err == nil {
	// 				// 删除成功，记录日志并返回
	// 				log.Printf("nats: Successfully deleted cache key %s on attempt %d", msg.Key, attempt+1)
	// 				return
	// 			} else {
	// 				lastErr = err
	// 				// 如果不是最后一次重试，等待后继续
	// 				if attempt < maxRetryAttempts-1 {
	// 					// 使用指数退避策略，每次重试的等待时间翻倍
	// 					backoff := retryDelay * time.Duration(1<<uint(attempt)) // 2, 4, 8
	// 					log.Printf("Failed to delete cache key %s on attempt %d: %v, retrying in %v",
	// 						msg.Key, attempt+1, err, backoff)
	// 					time.Sleep(backoff)
	// 				}
	// 			}
	// 		}

	// 		// 达到最大重试次数后仍然失败
	// 		log.Printf("nats: Failed to delete cache key %s after %d attempts, last error: %v",
	// 			msg.Key, maxRetryAttempts, lastErr)
	// 	}
	// })
}

func (p *NatsCacheProcessor) Close() {
	p.nc.Close()
}

// AddDeleteCacheMessage 添加删除缓存的消息到消息队列
// key: 缓存的 key
// Subject: cache.product
// Operation: delete_list
// Attempts: 0
// MaxAttempts: 3
func (p *NatsCacheProcessor) AddDeleteCacheMessage(key string) error {
	// msg := &CacheMessage{
	// 	Key:       key,
	// 	Operation: messageQueueDeleteListKey,
	// 	Attempts:  0,
	// 	CreatedAt: time.Now(),
	// }

	// msgBytes, err := json.Marshal(msg)
	// if err != nil {
	// 	return err
	// }
	// return p.nc.Publish(cacheProductSubject, msgBytes)
	return nil
}
