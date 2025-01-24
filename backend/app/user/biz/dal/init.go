package dal

import (
	"github.com/tiktokmall/backend/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
