package mq

import (
	"github.com/tiktokmall/backend/app/user/biz/dal/redis"
	"github.com/tiktokmall/backend/app/user/conf"

	"github.com/nats-io/nats.go"
)

var (
	Nc     *nats.Conn
	err    error
	Natscp *NatsCacheProcessor
)

func Init() {
	// 需要提前初始化 redis
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
	}
	Natscp = InitNatsCacheProcessor(Nc, redis.RedisClient)
	Natscp.StartProcessor()
}
