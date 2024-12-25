package BusinessConfig

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/initialize"
	interfaces "github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/xue-ding-e/gin"
)

var _ interfaces.Plugin = (*plugin)(nil)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) Register(group *gin.Engine) {
	ctx := context.Background()
	initialize.Iinit(ctx)
	initialize.Api(ctx)
	initialize.Router(group)
	initialize.Menu(ctx)
}
