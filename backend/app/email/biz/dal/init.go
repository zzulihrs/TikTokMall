package dal

import (
	"github.com/tiktokmall/backend/app/email/biz/dal/mysql"
	"github.com/tiktokmall/backend/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
