package main

import (
	"net/http"
	"strconv"

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

// 做什么：GET /api/v1/products/:id —— 取 id → 查商品 → 按结果给 400/404/500/200
func (h *ProductHandler) Get(c *gin.Context) {
	idStr := c.Param("id")                      // 路径里的那一段，例如 "1" / "999" / "abc"
	id, err := strconv.ParseUint(idStr, 10, 64) // 字符串转 uint64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id 必须是数字"}) // 客户端传错 → 400
		return
	}

	product, found, err := h.repo.GetProductByID(c.Request.Context(), id) // 三样东西一起接
	if err != nil {                                                       // 服务端故障 → 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found { // 查得到、但没这条数据 → 404
		c.JSON(http.StatusNotFound, gin.H{"error": "商品不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product}) // 成功 → 200
}
