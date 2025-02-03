package dal

import (
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
