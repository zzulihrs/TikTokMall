package dal

import (
	"github.com/tiktokmall/backend/app/checkout/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
