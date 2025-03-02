package model

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 商品详情暂时不做缓存
// 目前，只对商品列表做缓存
// 商品列表的缓存策略：
// 1. 缓存过期：为了避免缓存雪崩，每个缓存项都设置一个随机的过期时间
// 2. 缓存穿透：当请求一个不存在的商品列表时，返回空值，并设置一个较短的过期时间
// 3. 缓存击穿：使用互斥锁，只允许一个请求进入数据库查询，其他请求等待
// 4. 缓存更新：当商品信息发生变化时，删除缓存。店家服务会更新商品信息，同时删除商品列表缓存。

type ProductDAO struct {
	ctx   context.Context
	db    *gorm.DB
	redis *redis.Client
	locks *sync.Map // 用于缓存击穿防护的互斥锁
	// bloomFilter *BloomFilter // 布隆过滤器
}

func NewProductDAO(ctx context.Context, db *gorm.DB, redis *redis.Client) *ProductDAO {
	return &ProductDAO{
		ctx:   ctx,
		db:    db,
		redis: redis,
		locks: GetProductLocks(),
	}
}

/*
带缓存的商品列表。返回的商品包含以下信息。

"id": 3,
"name": "T-Shirt",
"description": "The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.",
"picture": "/static/image/t-shirt.jpeg",
"price": 6.6,
"ownerId": 1,
"ownerName": "root"
*/

type ProductListItem struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	OwnerId     int64   `json:"owner_id"`
	OwnerName   string  `json:"owner_name"`
}

func (dao *ProductDAO) GetProductInfoListByQuery(query string) (products []*ProductListItem, err error) {
	// 从缓存中获取
	cachedKey := getProductListCacheKey(query, "")
	// 第一次缓存检查
	val, err := dao.redis.Get(dao.ctx, cachedKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("product not found (cached)")
		}
		err = json.Unmarshal([]byte(val), &products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}

	// 缓存未命中，从数据库中获取
	// 缓存击穿防护
	mu, _ := dao.locks.LoadOrStore(cachedKey, &sync.Mutex{})
	lock := mu.(*sync.Mutex)

	// 使用TryLock实现带超时的锁获取
	start := time.Now()
	// TODO: 可以考虑时用 channel，需要注释协程是否正确关闭
	for {
		if lock.TryLock() {
			defer lock.Unlock() // 确保最终解锁
			break
		}
		if time.Since(start) > 10*time.Second {
			return nil, fmt.Errorf("系统繁忙，请稍后重试")
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 再次检查缓存是否存在
	val, err = dao.redis.Get(dao.ctx, cachedKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("product not found (cached)")
		}
		err = json.Unmarshal([]byte(val), &products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}
	// 缓存未命中，从数据库中获取
	// 从数据库中获取
	ps := []*Product{}
	err = dao.db.WithContext(dao.ctx).Model(&Product{}).
		Find(&ps, "name like ? or description like ? and deleted_at IS NULL", "%"+query+"%", "%"+query+"%").
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 空值查询，缓存空值防止缓存穿透，随机过期时间防雪崩
			dao.redis.Set(dao.ctx, cachedKey, "", getCachedExpiration(emptyCacheDuration, emptyCacheDuration))
		}
		return nil, err
	}

	products = make([]*ProductListItem, len(ps))

	merchantSet := make(map[int]int)
	merchantIDs := []int{}

	for _, p := range ps {
		if merchantSet[p.MerchantID] == 0 {
			merchantSet[p.MerchantID] = len(merchantIDs) + 1
			merchantIDs = append(merchantIDs, p.MerchantID)
		}
	}
	merchants := []Merchant{}
	err = dao.db.WithContext(dao.ctx).Model(&Merchant{}).Find(&merchants, "id in (?) AND deleted_at IS NULL", merchantIDs).Error
	if err != nil {
		return nil, err
	}

	for i, p := range ps {
		products[i] = &ProductListItem{
			ID:          int64(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			OwnerId:     int64(p.MerchantID),
			OwnerName:   merchants[merchantSet[p.MerchantID]-1].Username,
		}
	}

	// 缓存结果
	if data, err := json.Marshal(products); err == nil {
		dao.redis.Set(dao.ctx, cachedKey, string(data), getCachedExpiration(cacheDuration, cacheDuration))
	}

	return products, nil
}

func (dao *ProductDAO) GetProductInfoListByCategory(category string) (products []*ProductListItem, err error) {
	// 从缓存中获取
	cachedKey := getProductListCacheKey("", category)
	// 第一次缓存检查
	val, err := dao.redis.Get(dao.ctx, cachedKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("product not found (cached)")
		}
		err = json.Unmarshal([]byte(val), &products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}

	// 缓存未命中，从数据库中获取
	// 缓存击穿防护
	mu, _ := dao.locks.LoadOrStore(cachedKey, &sync.Mutex{})
	lock := mu.(*sync.Mutex)

	// 使用TryLock实现带超时的锁获取
	start := time.Now()
	// TODO: 可以考虑时用 channel，需要注释协程是否正确关闭
	for {
		if lock.TryLock() {
			defer lock.Unlock() // 确保最终解锁
			break
		}
		if time.Since(start) > 10*time.Second {
			return nil, fmt.Errorf("系统繁忙，请稍后重试")
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 再次检查缓存是否存在
	val, err = dao.redis.Get(dao.ctx, cachedKey).Result()
	if err == nil {
		if val == "" { // 空值表示缓存穿透保护
			return nil, fmt.Errorf("product not found (cached)")
		}
		err = json.Unmarshal([]byte(val), &products)
		if err != nil {
			return nil, err
		}
		return products, nil
	}
	// 缓存未命中，从数据库中获取
	// 从数据库中获取
	cs := []*Category{}
	err = dao.db.WithContext(dao.ctx).Model(Category{}).Where(&Category{Name: category}).
		Preload("Products").Find(&cs, "deleted_at IS NULL").
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 空值查询，缓存空值防止缓存穿透，随机过期时间防雪崩
			dao.redis.Set(dao.ctx, cachedKey, "", getCachedExpiration(emptyCacheDuration, emptyCacheDuration))
		}
		return nil, err
	}

	merchantSet := make(map[int]int)
	merchantIDs := []int{}

	products = make([]*ProductListItem, 0)
	for _, c := range cs {
		for _, p := range c.Products {
			products = append(products, &ProductListItem{
				ID:          int64(p.ID),
				Name:        p.Name,
				Price:       p.Price,
				Description: p.Description,
				Picture:     p.Picture,
				OwnerId:     int64(p.MerchantID),
			})
			if merchantSet[p.MerchantID] == 0 {
				merchantSet[p.MerchantID] = len(merchantIDs) + 1
				merchantIDs = append(merchantIDs, p.MerchantID)
			}
		}
	}
	merchants := []Merchant{}
	err = dao.db.WithContext(dao.ctx).Model(&Merchant{}).Find(&merchants, "id in (?) and deleted_at IS NULL", merchantIDs).Error
	if err != nil {
		return nil, err
	}

	for i, p := range products {
		products[i] = &ProductListItem{
			OwnerName: merchants[merchantSet[int(p.OwnerId)]-1].Username,
		}
	}

	// 缓存结果
	if data, err := json.Marshal(products); err == nil {
		dao.redis.Set(dao.ctx, cachedKey, string(data), getCachedExpiration(cacheDuration, cacheDuration))
	}

	return products, nil
}

type ProductDetail struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Picture        string   `json:"picture"`
	SliderImgs     string   `json:"slider_imgs"`
	Price          float32  `json:"price"`
	Stock          int32    `json:"stock"`
	OwnerId        int64    `json:"owner_id"`
	OwnerName      string   `json:"owner_name"`
	CategorieNames []string `json:"categories"`
}

func (dao *ProductDAO) GetProductDetailByPid(pid int64) (product *ProductDetail, err error) {
	// 从缓存中获取
	cachedKey := getProductDetailCacheKey(pid)
	err = getValueByKeyFromCache(dao.ctx, dao.redis, cachedKey, &product)
	if err == nil {
		return product, nil
	}
	if errors.Is(err, cacheEmptyErr) {
		return nil, fmt.Errorf("product not found (cached)")
	}

	// 缓存未命中，从数据库中获取
	// 缓存击穿防护
	mu, _ := dao.locks.LoadOrStore(cachedKey, &sync.Mutex{})
	lock := mu.(*sync.Mutex)

	// 使用TryLock实现带超时的锁获取
	start := time.Now()
	// TODO: 可以考虑时用 channel，需要注释协程是否正确关闭
	for {
		if lock.TryLock() {
			defer lock.Unlock() // 确保最终解锁
			break
		}
		if time.Since(start) > 10*time.Second {
			return nil, fmt.Errorf("系统繁忙，请稍后重试")
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 再次检查缓存是否存在
	err = getValueByKeyFromCache(dao.ctx, dao.redis, cachedKey, &product)
	if err == nil {
		return product, nil
	}
	if errors.Is(err, cacheEmptyErr) {
		return nil, fmt.Errorf("product not found (cached)")
	}

	// 从数据库中获取
	p := &Product{}
	err = dao.db.WithContext(dao.ctx).Model(&Product{}).Preload("Categories").Where("id = ? AND deleted_at IS NULL", pid).First(p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 空值查询，缓存空值防止缓存穿透，随机过期时间防雪崩
			dao.redis.Set(dao.ctx, cachedKey, "", getCachedExpiration(emptyCacheDuration, emptyCacheDuration))
		}
		return nil, err
	}

	merchant := &Merchant{}
	err = dao.db.WithContext(dao.ctx).Model(&Merchant{}).Where("id = ? AND deleted_at IS NULL", p.MerchantID).First(merchant).Error
	if err != nil {
		return nil, err
	}

	categoryNames := make([]string, len(p.Categories))
	for i, c := range p.Categories {
		categoryNames[i] = c.Name
	}

	product = &ProductDetail{
		ID:             int64(p.ID),
		Name:           p.Name,
		Description:    p.Description,
		Picture:        p.Picture,
		SliderImgs:     p.SliderImgs,
		Price:          p.Price,
		Stock:          p.Stock,
		OwnerId:        int64(p.MerchantID),
		OwnerName:      merchant.Username,
		CategorieNames: categoryNames,
	}

	// 缓存结果
	if data, err := json.Marshal(product); err == nil {
		dao.redis.Set(dao.ctx, cachedKey, string(data), getCachedExpiration(cacheDuration, cacheDuration))
	}

	return product, nil
}
