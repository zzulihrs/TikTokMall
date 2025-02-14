package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddCart(ctx context.Context, db *gorm.DB, cart *Cart) error {
	var row Cart

	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
		First(&row).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID != 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", cart.Qty)).Error
	}
	return db.WithContext(ctx).Model(&Cart{}).Create(cart).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Delete(&Cart{}).Error
}

func ChangeQty(ctx context.Context, db *gorm.DB, cart *Cart) error {
	// 先判断数字范围是否合格
	// 商品数量小于0----前端页面限制不能小于0
	// 商品数量为0---删除购物车中该商品
	if cart.Qty == 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
			Delete(&Cart{}).Error
	}
	// 商品数量合理
	return db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: cart.UserId, ProductId: cart.ProductId}).
		UpdateColumn("qty", cart.Qty).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) (cartList []*Cart, err error) {
	if db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err = db.Debug().
		WithContext(ctx).
		Model(&Cart{}).
		Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}
