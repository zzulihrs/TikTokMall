package model

import (
	"context"

	"gorm.io/gorm"
)

type Merchant struct {
	Base
	UserId   int    `json:"user_id" gorm:"unique_index"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique" gorm:"uniqueIndex:idx_email(255)"`
	// 添加反向关联
	Products []Product `gorm:"foreignKey:merchant_id"`
}

func (Merchant) TableName() string {
	return "merchant"
}

type MerchantQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (m *MerchantQuery) GetById(merchantId int) (merchant Merchant, err error) {
	err = m.db.WithContext(m.ctx).Model(&Merchant{}).First(&merchant, merchantId).Error
	return
}

func (m *MerchantQuery) GetManyById(ids []int) (merchants []*Merchant, err error) {
	err = m.db.WithContext(m.ctx).Model(&Merchant{}).Find(&merchants, "id in (?)", ids).Error
	return
}
func (m *MerchantQuery) InsertOne(merchant Merchant) (err error) {
	err = m.db.WithContext(m.ctx).Model(&Merchant{}).Create(&merchant).Error
	return
}
func (m *MerchantQuery) InsertMany(merchants []Merchant) (err error) {
	err = m.db.WithContext(m.ctx).Model(&Merchant{}).Create(&merchants).Error
	return
}

func NewMerchantQuery(ctx context.Context, db *gorm.DB) *MerchantQuery {
	return &MerchantQuery{
		ctx: ctx,
		db:  db,
	}
}
