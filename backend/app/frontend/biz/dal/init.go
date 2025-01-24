package dal

import (
	"github.com/tiktokmall/backend/app/frontend/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
