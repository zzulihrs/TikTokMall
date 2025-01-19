package dal

import (
	"bytetech/course/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
