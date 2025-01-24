package dal

import (
	"github.com/tiktokmall/backend/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
