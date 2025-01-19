package dal

import (
	"bytetech/course/app/email/biz/dal/mysql"
	"bytetech/course/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
