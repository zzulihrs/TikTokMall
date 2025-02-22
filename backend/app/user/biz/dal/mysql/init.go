package mysql

import (
	"fmt"
	"os"

	"github.com/tiktokmall/backend/app/user/biz/model"
	"github.com/tiktokmall/backend/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.User{})
		DB.AutoMigrate( // nolint: errcheck
			&model.User{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `user` (`id`,`created_at`,`updated_at`,`email`,`password_hashed`, `username`) VALUES (1,'2023-12-26 09:46:19.852','2023-12-26 09:46:19.852','root@example.com','$2a$10$gnw261yJ8zF1qyh/lBUnU.IX1w3j1xFA3sxf4QvlVpwfEU69ivuqK', 'root')")
		}
	}
}
