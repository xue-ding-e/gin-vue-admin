package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Sku = new(sku)

type sku struct{}

// Init 初始化 sku表 路由信息
func (self *sku) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("sku").Use(middleware.OperationRecord())
		group.POST("createSku", apiSku.CreateSku)             // 新建sku表
		group.DELETE("deleteSku", apiSku.DeleteSku)           // 删除sku表
		group.DELETE("deleteSkuByIds", apiSku.DeleteSkuByIds) // 批量删除sku表
		group.PUT("updateSku", apiSku.UpdateSku)              // 更新sku表
	}
	{
		group := private.Group("sku")
		group.GET("findSku", apiSku.FindSku)       // 根据ID获取sku表
		group.GET("getSkuList", apiSku.GetSkuList) // 获取sku表列表
	}
	{
		group := public.Group("sku")
		group.GET("getSkuPublic", apiSku.GetSkuPublic) // sku表开放接口
	}
}
