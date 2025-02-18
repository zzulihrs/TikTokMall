package model

import (
	"context"
	"log"
	"testing"
)

func TestProductQuery_SearchProducts(t *testing.T) {
	productQuery := NewProductQuery(context.Background(), testDB)
	product, err := productQuery.SearchProducts("test")
	log.Println(product, err)

}
