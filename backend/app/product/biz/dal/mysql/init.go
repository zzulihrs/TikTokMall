package mysql

import (
	"fmt"
	"os"

	"github.com/tiktokmall/backend/app/product/biz/model"
	"github.com/tiktokmall/backend/app/product/conf"

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
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Merchant{},
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			// 先插入 merchant 表数据
			DB.Exec("INSERT INTO `product`.`merchant` VALUES (1, '2023-12-06 15:05:06', '2023-12-06 15:05:06', NULL, 100, 'cloudwego', 'cloudwego@example.com'), (2, '2024-10-06 15:05:06', '2024-12-06 15:05:06', NULL, 101, 'test', 'test@example.com'), (3, '2024-12-06 15:05:06', '2024-12-06 15:05:06', NULL, 102, 'test', 'test@test.com')")
			// 再插入 category 表数据
			DB.Exec("INSERT INTO `product`.`category` VALUES (1, '2023-12-06 15:05:06','2023-12-06 15:05:06', NULL, 'T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06', NULL, 'Sticker','Sticker')")
			// 接着插入 product 表数据
			DB.Exec("INSERT INTO `product`.`product` VALUES (1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90, '', 100000, 1), (2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', NULL, 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80, '', 100000, 1), (3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60, '', 100000, 1), (4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20, '', 100000, 1), (5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', NULL, 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10, '', 100000, 1), (6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80, '', 100000, 1), (7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', NULL, 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80, '', 100000, 1)")

			// 最后插入 product_category 表数据
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		}
	}
	// if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
	// 	panic(err)
	// }
}
