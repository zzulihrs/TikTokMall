package consumer

import (
	"github.com/redis/go-redis/v9"
	"github.com/tiktokmall/backend/app/user/biz/consumer/user"
	userRedis "github.com/tiktokmall/backend/app/user/biz/dal/redis"
)

var (
	redisCli *redis.Client
)

func Init() {
	redisCli = userRedis.RedisClient
	user.ConsumerInit()
}
