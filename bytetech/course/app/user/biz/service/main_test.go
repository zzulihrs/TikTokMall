package service

import (
	dalMysql "bytetech/course/app/user/biz/dal/mysql"
	"bytetech/course/app/user/biz/model"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	var err error
	dalMysql.DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !dalMysql.DB.Migrator().HasTable(&model.User{})
		dalMysql.DB.AutoMigrate( // nolint: errcheck
			&model.User{},
		)
		if needDemoData {
			dalMysql.DB.Exec("INSERT INTO `user` (`id`,`created_at`,`updated_at`,`email`,`password_hashed`) VALUES (1,'2023-12-26 09:46:19.852','2023-12-26 09:46:19.852','123@admin.com','$2a$10$jTvUFh7Z8Kw0hLV8WrAws.PRQTeuH4gopJ7ZMoiFvwhhz5Vw.bj7C')")
		}
	}

	os.Exit(m.Run())
}
