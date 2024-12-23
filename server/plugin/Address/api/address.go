package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Address/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Address/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Address = new(AddressApi)

type AddressApi struct{}

// CreateAddress 创建地址
// @Tags Address
// @Summary 创建地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "创建地址"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /address/createAddress [post]
func (self *AddressApi) CreateAddress(c *fiber.Ctx) {
	var address model.Address
	err := c.BodyParser(&address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		address.UserID = utils.GetUserID(c)
	}
	err = addressService.CreateAddress(&address)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAddress 删除地址
// @Tags Address
// @Summary 删除地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "删除地址"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /address/deleteAddress [delete]
func (self *AddressApi) DeleteAddress(c *fiber.Ctx) {
	ID := c.Query("ID")
	var UserID uint
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		UserID = utils.GetUserID(c)
	}
	err := addressService.DeleteAddress(ID, UserID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAddressByIds 批量删除地址
// @Tags Address
// @Summary 批量删除地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /address/deleteAddressByIds [delete]
func (self *AddressApi) DeleteAddressByIds(c *fiber.Ctx) {
	IDs := c.QueryArray("ID")
	var UserID uint
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		UserID = utils.GetUserID(c)
	}
	if err := addressService.DeleteAddressByIds(IDs, UserID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAddress 更新地址
// @Tags Address
// @Summary 更新地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Address true "更新地址"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /address/updateAddress [put]
func (self *AddressApi) UpdateAddress(c *fiber.Ctx) {
	var address model.Address
	err := c.BodyParser(&address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		address.UserID = utils.GetUserID(c)
	}
	err = addressService.UpdateAddress(address)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAddress 用id查询地址
// @Tags Address
// @Summary 用id查询地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Address true "用id查询地址"
// @Success 200 {object} response.Response{data=model.Address,msg=string} "查询成功"
// @Router /address/findAddress [get]
func (self *AddressApi) FindAddress(c *fiber.Ctx) {
	ID := c.Query("ID")
	var UserID uint
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		UserID = utils.GetUserID(c)
	}
	if readdress, err := addressService.GetAddress(ID, UserID, authID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"readdress": readdress}, c)
	}
}

// getDefaultAddress 获取用户默认地址
// @Tags Address
// @Summary 获取用户默认地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/getDefaultAddress [get]
func (self *AddressApi) GetDefaultAddress(c *fiber.Ctx) {
	var UserID uint
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		UserID = utils.GetUserID(c)
	}

	address, err := addressService.GetDefaultAddress(UserID) // 接收两个返回值
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(address, c) // 使用返回的地址数据
	}
}

// GetAddressList 分页获取地址列表
// @Tags Address
// @Summary 分页获取地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.AddressSearch true "分页获取地址列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /address/getAddressList [get]
func (self *AddressApi) GetAddressList(c *fiber.Ctx) {
	var pageInfo request.AddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authID := utils.GetUserAuthorityId(c)
	if authID != 888 {
		pageInfo.UserID = utils.GetUserID(c)
	}
	list, total, err := addressService.GetAddressInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAddressDataSource 获取Address的数据源
// @Tags Address
// @Summary 获取Address的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /address/getAddressDataSource [get]
func (self *AddressApi) GetAddressDataSource(c *fiber.Ctx) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := addressService.GetAddressDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}
