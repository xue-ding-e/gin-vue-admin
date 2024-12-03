package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "businessConfig/updateBusinessConfig",
			Description: "更新业务设置",
			ApiGroup:    "业务设置",
			Method:      "POST",
		}, {
			Path:        "businessConfig/getBusinessConfig",
			Description: "获取业务设置",
			ApiGroup:    "业务设置",
			Method:      "GET",
		}, {
			Path:        "businessConfig/refeshConfigCacheToEtcd",
			Description: "刷新业务设置到etcd",
			ApiGroup:    "业务设置",
			Method:      "GET",
		}, {
			Path:        "businessConfig/refeshConfigCacheToDB",
			Description: "刷新业务设置缓存到数据库",
			ApiGroup:    "业务设置",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
