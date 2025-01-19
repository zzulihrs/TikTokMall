package dal

import (
	"bytetech/course/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
