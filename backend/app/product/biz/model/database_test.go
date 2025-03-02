package model

import (
	"context"
	"log"
	"testing"

	"github.com/tiktokmall/backend/app/product/biz/util"
)

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
