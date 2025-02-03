package dal

import (
	"github.com/tiktokmall/backend/app/payment/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
