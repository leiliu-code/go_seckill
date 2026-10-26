package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo *ProductRepo
} //另一个结构体，用来处理http请求

func NewProductHandler(repo *ProductRepo) *ProductHandler {
	return &ProductHandler{repo: repo}
} //返回新指针，依赖repo层

func (h *ProductHandler) List(c *gin.Context) {
	list, err := h.repo.ListProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": list})
} /*处理http的主函数，传入c这个结构体指针*gin.Context，引用Request这个字段，调用 Request 的 Context 方法，拿到本次请求的 context.Context
返回一个context.Context类型的对象，然后把这个对象传入到repo里面的ListProducts函数里面，
返回商品列表，最后把这个商品信息转为json格式*/
