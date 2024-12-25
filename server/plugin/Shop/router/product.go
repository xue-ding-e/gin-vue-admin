package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Product = new(product)

type product struct{}

// Init 初始化 产品 路由信息
func (self *product) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("product").Use(middleware.OperationRecord())
		group.POST("createProduct", apiProduct.CreateProduct)             // 新建产品
		group.DELETE("deleteProduct", apiProduct.DeleteProduct)           // 删除产品
		group.DELETE("deleteProductByIds", apiProduct.DeleteProductByIds) // 批量删除产品
		group.PUT("updateProduct", apiProduct.UpdateProduct)              // 更新产品
	}
	{
		group := private.Group("product")
		group.GET("findProduct", apiProduct.FindProduct)       // 根据ID获取产品
		group.GET("getProductList", apiProduct.GetProductList) // 获取产品列表
	}
	{
		//group := public.Group("product")
		//group.GET("getProductPublic", apiProduct.GetProductPublic)  // 产品开放接口
	}
}
