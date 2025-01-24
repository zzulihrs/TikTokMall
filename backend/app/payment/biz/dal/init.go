package dal

import (
	"github.com/tiktokmall/backend/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
