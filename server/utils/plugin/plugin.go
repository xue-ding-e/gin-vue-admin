package plugin

import (
	"github.com/xue-ding-e/gin"
)

const (
	OnlyFuncName = "Plugin"
)

// Plugin 插件模式接口化
type Plugin interface {
	// Register 注册路由
	Register(group *gin.RouterGroup)

	// RouterPath 用户返回注册路由
	RouterPath() string
}
