package model

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tiktokmall/backend/app/product/biz/util"
	"github.com/tiktokmall/backend/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	os.Chdir("../../")
	_ = godotenv.Load()
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	var err error
	testDB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestAddMerchant(t *testing.T) {
	var merchants []Merchant
	for i := 0; i < 10; i++ {
		merchants = append(merchants, Merchant{
			Email:    util.GenerateRandomEmail(10),
			Username: util.GenerateRandomString(5),
		})
	}
	log.Println(merchants)
	err := NewMerchantQuery(context.TODO(), testDB).InsertMany(merchants)
	log.Println(err)
}

func TestInitCategory(t *testing.T) {
	categories := []Category{
		{
			Name:        "T-Shirt",
			Description: "T-Shirt",
		},
		{
			Name:        "Sticker",
			Description: "Sticker",
		},
		{
			Name:        "家居用品",
			Description: "家居用品",
		},
		{
			Name:        "食品",
			Description: "食品",
		},
		{
			Name:        "汽车用品",
			Description: "汽车用品",
		},
		{
			Name:        "其他",
			Description: "其他",
		},
	}
	query := NewCategoryQuery(context.TODO(), testDB)
	err := query.InsertMany(categories)
	log.Println(err)
}

func TestInitProduct(t *testing.T) {
	productQuery := NewProductQuery(context.TODO(), testDB)
	var products []Product

	productQuery.InsertMany(products)
}
