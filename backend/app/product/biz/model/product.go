package model

import (
	"context"
	"fmt"
	"log"
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
func (p *ProductQuery) InsertMany(products []Product) (err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Create(&products).Error
	return
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
	err1 := func() error {
		err2 := cachedResult.Err()
		if err2 != nil {
			return err2
		}
		cachedResultByte, err1 := cachedResult.Bytes()
		if err1 != nil {
			return err2
		}
		if len(cachedResultByte) == 0 {
			err = fmt.Errorf("product not found (cached)")
			return nil
		}
		err2 = json.Unmarshal(cachedResultByte, &product)
		if err2 != nil {
			return err2
		}
		return nil
	}()

	// 缓存命中，直接返回
	if err1 == nil {
		return product, err
	}

	// 缺少缓存击穿
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

	return
}

func NewCachedProductQuery(pq *ProductQuery, cachedClient *redis.Client, prefix string) CachedProductQuery {
	return CachedProductQuery{
		productQuery: pq,
		cacheClient:  cachedClient,
		prefix:       prefix,
	}
}
