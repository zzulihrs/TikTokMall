package dal

import (
	"github.com/tiktokmall/backend/app/merchant/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/merchant/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
