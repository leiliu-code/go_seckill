package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
} //在repo这个存储着数据库一些信息的文件里面把句柄作为了这个结构体重定义

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
} //将句柄传入到repo里面，返回一个新的指针

func (r *ProductRepo) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	err := r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		panic(fmt.Errorf("func ListProducts,table products, err: %w", err))
	}
	return products, err
} //先定义了一个空商品切片，WithContext对句柄进行包装，Find函数查询数据库里面的products表，把结果放到products切片里面，最后返回这个切片和错误信息
