package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{
			Path:        "/minio/getMinioSts",
			Description: "获取minio sts凭证",
			ApiGroup:    "minio",
			Method:      "GET",
		},
	}
	utils.RegisterApis(entities...)
}
