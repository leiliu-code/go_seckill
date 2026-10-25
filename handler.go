package main

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo *ProductRepo
} //另一个结构体，用来处理http请求

func NewProductHandler(repo *ProductRepo) *ProductHandler {
	return &ProductHandler{repo: repo}
} //返回新指针，在不同部分新定义商品，知道对商品做了什么

func (h *ProductHandler) List(c *gin.Context) {
	list, err := h.repo.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"products": list})
} /*处理http的主函数，通过句柄获取Context一些信息，
  然后调用repo里面的ListProducts函数，
  返回商品列表，最后把这个商品信息转为json格式*/
