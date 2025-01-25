package dal

import (
	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
