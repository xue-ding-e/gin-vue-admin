package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{ApiGroup: "用户地址", Method: "POST", Path: "/address/createAddress", Description: "新增用户地址"},
		{ApiGroup: "用户地址", Method: "DELETE", Path: "/address/deleteAddress", Description: "删除用户地址"},
		{ApiGroup: "用户地址", Method: "DELETE", Path: "/address/deleteAddressByIds", Description: "批量删除用户地址"},
		{ApiGroup: "用户地址", Method: "PUT", Path: "/address/updateAddress", Description: "更新用户地址"},
		{ApiGroup: "用户地址", Method: "GET", Path: "/address/findAddress", Description: "根据ID获取用户地址"},
		{ApiGroup: "用户地址", Method: "GET", Path: "/address/getAddressList", Description: "获取用户地址列表"},
		{ApiGroup: "用户地址", Method: "GET", Path: "/address/getDefaultAddress", Description: "获取默认用户地址"},
	}
	utils.RegisterApis(entities)
}
