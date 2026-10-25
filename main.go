package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")) //定义一个dsn字符串，里面包含了数据库的用户名，密码，主机，端口和数据库名，不懂

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //不懂
	if err != nil {
		panic(fmt.Errorf("connect mysql: %w", err))
	} //如果有错误信息，程序崩溃报错
	repo := NewProductRepo(db)         //不懂db怎么来的
	handler := NewProductHandler(repo) //重定义handler，传入repo这个指针，方便知道处理逻辑

	r := gin.Default()                      //不懂
	r.GET("/api/v1/products", handler.List) //写路由路径，传入商品信息
	_ = r.Run(":8080")                      //在常用端口上面运行这个程序
}
