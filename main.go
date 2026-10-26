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
		"dev", //DB_USER
		os.Getenv("DB_PASS"),
		"127.0.0.1", //DB_HOST
		"3306",      //DB_PORT
		"seckill")   //DB_NAME
	/*定义一个dsn字符串，里面包含了数据库的用户名，密码，主机，端口和数据库名，
	把要发送给数据库的参数写出来，设置gorm链接字符集为utf8mb4，parsetime把默认的false改为true，驱动把它解析成 time.Time，loc获取本地时区
	从进程的环境变量中拿相关信息，设置GORM的参数*/

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //调返回 *gorm.DB 句柄，内部已建连接池并连过库
	if err != nil {
		panic(fmt.Errorf("connect mysql: %w", err))
	} //如果有错误信息，程序崩溃报错
	repo := NewProductRepo(db)         //注入依赖，repo依赖db
	handler := NewProductHandler(repo) //注入依赖，handler依赖repo

	r := gin.Default()                         //创建一个gin的默认实例，引擎engine，自带日志和恢复中间件
	r.GET("/api/v1/products", handler.List)    //路由，获取商品信息
	r.GET("/api/v1/products/:id", handler.Get) //路由，获取请求的单个商品信息

	err = r.Run(":8080") //在常用端口上面运行这个程序
	if err != nil {
		panic(fmt.Errorf("run server: %w", err))
	}
}
