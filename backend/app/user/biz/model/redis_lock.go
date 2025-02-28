package model

import (
	"sync"
	"time"
)

const (
// 互斥锁超时时间
// lockTimeout = 3 * time.Second
)

var (
	UserLocks     *sync.Map // 全局锁的实例，缓存击穿时，只放行一个进入 MYSQL
	onceUserLocks sync.Once
)

func GetUserLocks() *sync.Map {
	onceUserLocks.Do(func() {
		UserLocks = &sync.Map{}
	})
	return UserLocks
}

func getUserEmailKey(email string) string {
	return "user_email:" + email
}

func getCachedExpiration(baseDuration, randomDuration time.Duration) time.Duration {
	// 生成一个随机的时间偏移量
	offset := time.Duration(randomDuration) * time.Second
	// 返回基础时间加上随机偏移量的结果
	return baseDuration + offset
}
