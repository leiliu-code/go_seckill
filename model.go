package main

import (
	"time"
)

type OrderStatus int8

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusPaid
	OrderStatusCanceled
)

// Product 商品表模型，一行是一个可售商品
type Product struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 主键，自增
	Name      string    `gorm:"column:name" json:"name"`                      // 商品名称
	Price     int64     `gorm:"column:price" json:"price"`                    // 单价 分
	Stock     int32     `gorm:"column:stock" json:"stock"`                    // 剩余库存 件
	Version   int32     `gorm:"column:version" json:"version"`                // 乐观锁版本号，每次扣库存,预留字段，本项目未使用
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`          // 创建时间，GORM 自动写入；DB 默认值兜底
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`          // 最后更新时间，GORM 自动写入；DB 默认值兜底
}

// Order 订单表模型，一行是一次下单
type Order struct {
	ID        uint64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 主键 自增
	UserID    uint64      `gorm:"column:user_id" json:"user_id"`                // 下单用户标识符
	ProductID uint64      `gorm:"column:product_id" json:"product_id"`          // 购买商品标识符
	Quantity  int32       `gorm:"column:quantity" json:"quantity"`              // 购买数量 >=1
	Amount    int64       `gorm:"column:amount" json:"amount"`                  // 订单总金额 分
	Status    OrderStatus `gorm:"column:status" json:"status"`                  // 订单状态  0待支付 1已支付 2已取消
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at"`          // 下单时间 GORM 自动写入；DB 默认值兜底
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at"`          // 状态变更时间 GORM 自动写入；DB 默认值兜底
}
