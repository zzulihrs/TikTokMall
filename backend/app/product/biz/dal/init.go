package dal

import (
	"github.com/tiktokmall/backend/app/product/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
