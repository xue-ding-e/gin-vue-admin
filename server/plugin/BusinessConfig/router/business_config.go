package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/xue-ding-e/gin"
)

var BusinessConfig = new(businessConfig)

type businessConfig struct{}

// Init 初始化 业务设置 路由信息
func (r *businessConfig) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("businessConfig").Use(middleware.OperationRecord())
		group.POST("updateBusinessConfig", apiBusinessConfig.UpdateBusinessConfig)
	}
	{
		group := private.Group("businessConfig")
		group.GET("getBusinessConfig", apiBusinessConfig.GetBusinessConfig)
		//group.GET("refeshConfigCacheToEtcd", apiBusinessConfig.RefeshConfigCacheToEtcd)
		group.GET("refeshConfigCacheToDB", apiBusinessConfig.RefeshConfigCacheToDB)
	}
}
