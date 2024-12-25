package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/xue-ding-e/gin"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *gin.Engine) {
	PluginInitV2(engine, BusinessConfig.Plugin)
	PluginInitV2(engine, announcement.Plugin)
}
