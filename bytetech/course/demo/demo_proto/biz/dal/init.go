package dal

import (
	"bytetech/course/demo/demo_proto/biz/dal/mysql"
	"bytetech/course/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
