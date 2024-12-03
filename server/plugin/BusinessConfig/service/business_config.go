package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/model"
	config "github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/plugin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var BusinessConfig = new(businessConfig)

type businessConfig struct{}

// TODO 做全局并发锁控制还有原子性
func (s *businessConfig) UpdateBusinessConfig(reqrest *config.GLOBAL_CONFIG_TYPE) error {
	// 校验更新后的配置
	if err := reqrest.RequestValidateBusinessConfig(); err != nil {
		return err
	}
	// 更新缓存中的数据
	if err := copier.CopyWithOption(&config.GLOBAL_CONFIG, &reqrest, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		global.GVA_LOG.Error("更新缓存中的数据失败", zap.Error(err))
		return err
	}

	configMap := utils.StructToMap(config.GLOBAL_CONFIG)

	for key, value := range configMap {
		if err := global.GVA_DB.Model(&model.BusinessConfig{}).
			Where("name = ?", key).
			Update("value", value).Error; err != nil {
			global.GVA_LOG.Error("更新数据库中的数据失败", zap.Error(err))
			return err
		}
	}
	return nil
}
func (s *businessConfig) RefeshConfigCacheToDB() error {
	configMap := utils.StructToMap(config.GLOBAL_CONFIG)
	for key, value := range configMap {
		if err := global.GVA_DB.Model(&model.BusinessConfig{}).
			Where("name = ?", key).
			Update("value", value).Error; err != nil {
			global.GVA_LOG.Error("更新数据库中的数据失败", zap.Error(err))
			return err
		}
	}
	return nil
}
