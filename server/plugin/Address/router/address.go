package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Address = new(address)

type address struct{}

// Init 初始化 个人收货地址 路由信息
func (r *address) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	addressRouter := private.Group("address").Use(middleware.OperationRecord())
	addressRouterWithoutRecord := private.Group("address")
	addressRouterWithoutAuth := public.Group("address")
	{
		addressRouter.POST("createAddress", addressApi.CreateAddress)             // 新建地址
		addressRouter.DELETE("deleteAddress", addressApi.DeleteAddress)           // 删除地址
		addressRouter.DELETE("deleteAddressByIds", addressApi.DeleteAddressByIds) // 批量删除地址
		addressRouter.PUT("updateAddress", addressApi.UpdateAddress)              // 更新地址
	}
	{
		addressRouterWithoutRecord.GET("getDefaultAddress", addressApi.GetDefaultAddress) // 获取用户默认地址
		addressRouterWithoutRecord.GET("findAddress", addressApi.FindAddress)             // 根据ID获取地址
		addressRouterWithoutRecord.GET("getAddressList", addressApi.GetAddressList)       // 获取地址列表
	}
	{
		addressRouterWithoutAuth.GET("getAddressDataSource", addressApi.GetAddressDataSource) // 获取用户地址数据源
		//addressRouterWithoutAuth.GET("getAddressPublic", addressApi.GetAddressPublic)         // 获取用户地址列表
	}
}
