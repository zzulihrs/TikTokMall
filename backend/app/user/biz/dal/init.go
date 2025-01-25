package dal

import (
	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
