package dal

import (
	"bytetech/course/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
