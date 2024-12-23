package api

import (
	config "github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/service"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

var BusinessConfig = new(businessConfig)

type businessConfig struct{}

func (self *businessConfig) GetBusinessConfig(c *fiber.Ctx) {
	response.OkWithData(config.GLOBAL_CONFIG, c)
}

func (self *businessConfig) UpdateBusinessConfig(c *fiber.Ctx) {
	var reqrest config.GLOBAL_CONFIG_TYPE
	if err := c.BodyParser(&reqrest); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := reqrest.RequestValidateBusinessConfig(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 将更新的配置写入 数据库
	if err := service.BusinessConfig.UpdateBusinessConfig(&reqrest); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(reqrest, "更新成功", c)
}

func (self *businessConfig) RefeshConfigCacheToDB(c *fiber.Ctx) {
	if err := service.BusinessConfig.RefeshConfigCacheToDB(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

//func (self *businessConfig) RefeshConfigCacheToEtcd(c *fiber.Ctx) {
//	configBytes, err := json.Marshal(config.GLOBAL_CONFIG)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	var configMap map[string]interface{}
//	if err := json.Unmarshal(configBytes, &configMap); err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	for key, value := range configMap {
//		etcdKey := fmt.Sprintf("/business/%s", key)
//		valueStr := fmt.Sprintf("%v", value)
//		if _, err := etcd.ETCD_CLIENT.Put(context.TODO(), etcdKey, valueStr); err != nil {
//			response.FailWithMessage(fmt.Sprintf("无法写入 etcd: %s", err), c)
//			return
//		}
//	}
//	response.OkWithMessage("保存成功", c)
//}
