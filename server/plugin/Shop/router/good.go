package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Good = new(good)

type good struct{}

// Init 初始化 商品 路由信息
func (r *good) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("good").Use(middleware.OperationRecord())
		group.POST("createGood", apiGood.CreateGood)             // 新建商品
		group.DELETE("deleteGood", apiGood.DeleteGood)           // 删除商品
		group.DELETE("deleteGoodByIds", apiGood.DeleteGoodByIds) // 批量删除商品
		group.PUT("updateGood", apiGood.UpdateGood)              // 更新商品
	}
	{
		group := private.Group("good")
		group.GET("findGood", apiGood.FindGood)       // 根据ID获取商品
		group.GET("getGoodList", apiGood.GetGoodList) // 获取商品列表
	}
	{
		group := public.Group("good")
		group.GET("getGoodPublic", apiGood.GetGoodPublic) // 商品开放接口
	}
}
