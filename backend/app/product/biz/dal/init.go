package dal

import (
	"github.com/tiktokmall/backend/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
