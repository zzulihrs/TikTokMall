package dal

import (
	"bytetech/course/app/frontend/biz/dal/mysql"
	"bytetech/course/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
