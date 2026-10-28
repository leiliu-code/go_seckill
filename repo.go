package main

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 在repo这个存储着数据库一些信息的文件里面把句柄作为了这个结构体字段保存
type ProductRepo struct {
	db *gorm.DB
}

// 将句柄传入到repo里面，返回一个新的指针
func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// 先定义了一个空商品切片，WithContext对句柄进行包装，Find函数查询数据库里面的products表，把结果放到products切片里面，最后返回这个切片和错误信息
func (r *ProductRepo) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	err := r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("seckill databases,products table: %w", err) // 错误路径：nil + 包装后的错误
	}
	return products, nil
}

// 做什么：按主键查单个商品；返回 商品 / 找没找到 / 真错误 三样东西
func (r *ProductRepo) GetProductByID(ctx context.Context, id uint64) (Product, bool, error) {
	var product Product // 承接数据库返回的那一行
	// First 按主键查一条；查不到时 GORM 会返回固定的错误值 gorm.ErrRecordNotFound
	err := r.db.WithContext(ctx).First(&product, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) { // 是不是"没查到"这个特定的错
		return Product{}, false, nil // 没找到：不是故障，交给上层决定回 404
	}
	if err != nil { // 其它错误（库连不上、SQL 错、超时）才算真故障
		return Product{}, false, fmt.Errorf("get product by id %d: %w", id, err)
	}
	return product, true, nil // 找到了
}
