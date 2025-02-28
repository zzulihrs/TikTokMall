package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	OrderId       string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId        uint32      `gorm:"type:int(11)"`
	OrderStatus   uint32      `gorm:"type:int(11)`
	PaymentMethod string      `gorm:"type:varchar(100)"`
	UserCurrency  string      `gorm:"type:varchar(100)"`
	Consignee     Consignee   `gorm:"embedded"`
	OrderItem     []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItem").Find(&orders).Error
	return
}

func GetOrderStatus(db *gorm.DB, ctx context.Context, orderId string) (status uint32, err error) {
	err = db.WithContext(ctx).Model(&Order{}).Where(&Order{OrderId: orderId}).Select("order_status").Find(&status).Error
	return status, err
}

func ChangeOrderStatus(db *gorm.DB, ctx context.Context, orderId string, status uint32) error {
	// 使用GORM事务封装操作
	return db.Transaction(func(tx *gorm.DB) error {
		// 最佳实践：在事务中创建新的上下文副本
		// 防止外部上下文取消影响事务完整性
		txCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 订单状态不为0，说明订单状态已经修改，不能随意变更
		if currentStatus, err := GetOrderStatus(tx, txCtx, orderId); err != nil {
			return fmt.Errorf("get order status failed: %v", err)
		} else {
			if currentStatus != 0 {
				return fmt.Errorf("order %s already in target status %d", orderId, status)
			}
		}

		// 使用结构体验证确保字段映射正确
		// 建议改用明确WHERE条件避免零值陷阱
		query := tx.WithContext(txCtx).
			Model(&Order{}).
			Where("order_id = ?", orderId). // 明确查询条件
			Update("order_status", status)

		// 处理找不到记录的特殊情况
		if query.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("order %s not found", orderId)
		}

		return query.Error
	})
}
