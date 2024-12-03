package initialize

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/model"
	config "github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/plugin"
	"github.com/spf13/viper"
)

func Iinit(ctx context.Context) {
	v := viper.New()
	v.SetConfigName("config") // 配置文件名（无扩展名）
	v.SetConfigType("yaml")   // 配置文件类型
	v.AddConfigPath("业务设置/")  // 配置文件所在的目录

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}

	// 如果数据库未初始化，直接将默认配置解析到 GLOBAL_CONFIG 中
	if err := v.Unmarshal(&config.GLOBAL_CONFIG); err != nil {
		panic(fmt.Errorf("配置解析失败: %s \n", err))
	}

	if global.GVA_DB == nil {
		return
	}
	global.GVA_DB.AutoMigrate(
		model.BusinessConfig{},
	)
	tempViper := viper.New()
	var dbConfig []model.BusinessConfig
	global.GVA_DB.Find(&dbConfig)
	for _, v := range dbConfig {
		tempViper.Set(v.Name, v.Value)
	}
	if err := tempViper.Unmarshal(&config.GLOBAL_CONFIG); err != nil {
		panic(fmt.Errorf("配置解析失败: %s \n", err))
	}
	config.GLOBAL_CONFIG.ValidateBusinessConfig()
}
