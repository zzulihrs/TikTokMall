package mysql

import (
	"fmt"
	"os"

	"github.com/tiktokmall/backend/app/merchant/biz/model"
	"github.com/tiktokmall/backend/app/merchant/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // 自定义日志输出目的地，这里使用标准输出
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // 慢 SQL 阈值
	// 		LogLevel:                  logger.Info, // 日志级别
	// 		IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
	// 		Colorful:                  true,        // 启用彩色日志
	// 	},
	// )
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			// Logger:                 newLogger, // 测试用
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate( //nolint:errcheck
			&model.Merchant{},
			&model.Product{},
			&model.Category{},
		)

	}
}
