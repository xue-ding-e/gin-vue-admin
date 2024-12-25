package plugin

import (
	"github.com/xue-ding-e/gin"
)

// Plugin 插件模式接口化v2
type Plugin interface {
	// Register 注册路由
	Register(group *gin.Engine)
}
