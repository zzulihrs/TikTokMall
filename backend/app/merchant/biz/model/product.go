package model

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	SliderImgs  string  `json:"slider_imgs"`
	Stock       int32   `json:"stock"`
	// 添加外键字段
	MerchantID int `json:"merchant_id"`
	Merchant   Merchant
	// 商品和分类是多对多的关系，因此需要一个中间表来存储关系
	Categories []Category `json:"categories" gorm:"many2many:product_category;"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p *ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Preload("Categories").First(&product, productId).Where("deleted_at IS NULL").Error
	return
}

func (p *ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	log.Printf("q: %s", q)
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ? and deleted_at IS NULL",
		"%"+q+"%", "%"+q+"%").Error
	return
}
func (p *ProductQuery) GetProductListByCondition(condition string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where(condition).Find(&products).Error
	return
}
func (p *ProductQuery) InsertMany(products []Product) (err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Create(&products).Error
	return
}

func (p *ProductQuery) DeleteByPidAndMid(productId, merchantId int) (err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).UpdateColumn("deleted_at", time.Now()).Where("id = ? and merchant_id = ? and deleted_at IS NULL", productId, merchantId).Error
	return
}

func (p *ProductQuery) UpdateOne(product Product) (err error) {
	db := p.db.WithContext(p.ctx)
	var oldProduct Product
	err = db.Model(&Product{}).Find(&product, product.ID, product.MerchantID).First(oldProduct).Error
	if err != nil {
		return
	}
	updateFlag := false
	if product.Name != oldProduct.Name {
		oldProduct.Name = product.Name
		updateFlag = true
	}
	if product.Description != oldProduct.Description {
		oldProduct.Description = product.Description
		updateFlag = true
	}
	if product.Picture != oldProduct.Picture {
		oldProduct.Picture = product.Picture
		updateFlag = true
	}
	if product.Price != oldProduct.Price {
		oldProduct.Price = product.Price
		updateFlag = true
	}
	if product.Stock != oldProduct.Stock {
		oldProduct.Stock = product.Stock
		updateFlag = true
	}
	if product.SliderImgs != oldProduct.SliderImgs {
		oldProduct.SliderImgs = product.SliderImgs
		updateFlag = true
	}
	if !checkCategories(product.Categories, oldProduct.Categories) {
		oldProduct.Categories = product.Categories
		db.Model(&Product{}).Association("Categories").Clear()
	}
	if updateFlag {
		err = db.Model(&Product{}).Save(&oldProduct).Error
	}
	return
}
func checkCategories(cs1, cs2 []Category) bool {
	if len(cs1) != len(cs2) {
		return false
	}
	sort.Slice(cs1, func(i, j int) bool {
		return cs1[i].ID < cs1[j].ID
	})
	sort.Slice(cs2, func(i, j int) bool {
		return cs2[i].ID < cs2[j].ID
	})
	for i, c := range cs1 {
		if c.ID != cs2[i].ID {
			return false
		}
	}
	return true
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

type CachedProductQuery struct {
	productQuery *ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)
	// 解析从 redis 中获取的数据
	err = func() error {
		err1 := cachedResult.Err()
		if err1 != nil {
			return err1
		}
		cachedResultByte, err1 := cachedResult.Bytes()
		if err1 != nil {
			return err1
		}
		err1 = json.Unmarshal(cachedResultByte, &product)
		if err1 != nil {
			return err1
		}
		return nil
	}()

	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, nil
		}
		// json encode and cache the result
		encoded, err := json.Marshal(product)
		// encode 失败，直接返回商品信息
		if err != nil {
			return product, nil
		}
		// 缓存商品信息
		c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func NewCachedProductQuery(pq *ProductQuery, cachedClient *redis.Client, prefix string) CachedProductQuery {
	return CachedProductQuery{
		productQuery: pq,
		cacheClient:  cachedClient,
		prefix:       prefix,
	}
}
