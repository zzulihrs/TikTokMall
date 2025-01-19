package dal

import (
	"bytetech/course/demo/demo_thrift/biz/dal/mysql"
	"bytetech/course/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
