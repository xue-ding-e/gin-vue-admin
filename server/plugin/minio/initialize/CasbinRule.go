package initialize

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/pkg/errors"
)

func CasbinRule() error {
	if global.GVA_DB != nil {
		global.GVA_LOG.Debug("已存在数据库配置!")
		return nil
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "5555", V1: "/minio/getMinioSts", V2: "GET"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, "minio权限:Casbin 表 (casbin_rule) 数据初始化失败!")
	}
	return nil
}
