package dal

import (
	"github.com/tiktokmall/backend/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
