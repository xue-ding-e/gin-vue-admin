package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/xue-ding-e/gin"
)

var Brand = new(brand)

type brand struct{}

// Init 初始化 品牌表 路由信息
func (self *brand) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("brand").Use(middleware.OperationRecord())
		group.POST("createBrand", apiBrand.CreateBrand)             // 新建品牌表
		group.DELETE("deleteBrand", apiBrand.DeleteBrand)           // 删除品牌表
		group.DELETE("deleteBrandByIds", apiBrand.DeleteBrandByIds) // 批量删除品牌表
		group.PUT("updateBrand", apiBrand.UpdateBrand)              // 更新品牌表
	}
	{
		group := private.Group("brand")
		group.GET("findBrand", apiBrand.FindBrand)       // 根据ID获取品牌表
		group.GET("getBrandList", apiBrand.GetBrandList) // 获取品牌表列表
	}
	{
		group := public.Group("brand")
		group.GET("getBrandPublic", apiBrand.GetBrandPublic) // 品牌表开放接口
	}
}
