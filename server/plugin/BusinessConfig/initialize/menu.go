package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "businessConfig",
			Name:      "businessConfig",
			Hidden:    false,
			Component: "plugin/BusinessConfig/view/businessConfig.vue",
			Sort:      0,
			Meta:      model.Meta{Title: "业务设置"},
		},
	}
	utils.RegisterMenus(entities)
}
