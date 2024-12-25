package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Sku = new(sku)

type sku struct{}

// CreateSku 创建sku表
// @Tags Sku
// @Summary 创建sku表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sku true "创建sku表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sku/createSku [post]
func (self *sku) CreateSku(c *gin.Context) {
	var info model.Sku
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSku.CreateSku(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSku 删除sku表
// @Tags Sku
// @Summary 删除sku表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sku true "删除sku表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sku/deleteSku [delete]
func (self *sku) DeleteSku(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceSku.DeleteSku(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSkuByIds 批量删除sku表
// @Tags Sku
// @Summary 批量删除sku表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sku/deleteSkuByIds [delete]
func (self *sku) DeleteSkuByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := serviceSku.DeleteSkuByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSku 更新sku表
// @Tags Sku
// @Summary 更新sku表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sku true "更新sku表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sku/updateSku [put]
func (self *sku) UpdateSku(c *gin.Context) {
	var info model.Sku
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSku.UpdateSku(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSku 用id查询sku表
// @Tags Sku
// @Summary 用id查询sku表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Sku true "用id查询sku表"
// @Success 200 {object} response.Response{data=model.Sku,msg=string} "查询成功"
// @Router /sku/findSku [get]
func (self *sku) FindSku(c *gin.Context) {
	ID := c.Query("ID")
	resku, err := serviceSku.GetSku(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resku, c)
}

// GetSkuList 分页获取sku表列表
// @Tags Sku
// @Summary 分页获取sku表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SkuSearch true "分页获取sku表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sku/getSkuList [get]
func (self *sku) GetSkuList(c *gin.Context) {
	var pageInfo request.SkuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSku.GetSkuInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSkuPublic 不需要鉴权的sku表接口
// @Tags Sku
// @Summary 不需要鉴权的sku表接口
// @accept application/json
// @Produce application/json
// @Param data query request.SkuSearch true "分页获取sku表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sku/getSkuPublic [get]
func (self *sku) GetSkuPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceSku.GetSkuPublic()
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的sku表接口信息"}, "获取成功", c)
}
