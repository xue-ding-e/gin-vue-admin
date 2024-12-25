package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/model/request"
	"github.com/xue-ding-e/gin"
	"go.uber.org/zap"
)

var RichText = new(richText)

type richText struct{}

// CreateRichText 创建富文本
// @Tags RichText
// @Summary 创建富文本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RichText true "创建富文本"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /richText/createRichText [post]
func (a *richText) CreateRichText(c *gin.Context) {
	var info model.RichText
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceRichText.CreateRichText(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRichText 删除富文本
// @Tags RichText
// @Summary 删除富文本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RichText true "删除富文本"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /richText/deleteRichText [delete]
func (a *richText) DeleteRichText(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceRichText.DeleteRichText(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRichTextByIds 批量删除富文本
// @Tags RichText
// @Summary 批量删除富文本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /richText/deleteRichTextByIds [delete]
func (a *richText) DeleteRichTextByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := serviceRichText.DeleteRichTextByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRichText 更新富文本
// @Tags RichText
// @Summary 更新富文本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RichText true "更新富文本"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /richText/updateRichText [put]
func (a *richText) UpdateRichText(c *gin.Context) {
	var info model.RichText
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceRichText.UpdateRichText(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRichText 用id查询富文本
// @Tags RichText
// @Summary 用id查询富文本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RichText true "用id查询富文本"
// @Success 200 {object} response.Response{data=model.RichText,msg=string} "查询成功"
// @Router /richText/findRichText [get]
func (a *richText) FindRichText(c *gin.Context) {
	ID := c.Query("ID")
	rerichText, err := serviceRichText.GetRichText(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rerichText, c)
}

// GetRichTextList 分页获取富文本列表
// @Tags RichText
// @Summary 分页获取富文本列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.RichTextSearch true "分页获取富文本列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /richText/getRichTextList [get]
func (a *richText) GetRichTextList(c *gin.Context) {
	var pageInfo request.RichTextSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceRichText.GetRichTextInfoList(pageInfo)
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

// GetRichTextPublic 不需要鉴权的富文本接口
// @Tags RichText
// @Summary 不需要鉴权的富文本接口
// @accept application/json
// @Produce application/json
// @Param data query request.RichTextSearch true "分页获取富文本列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /richText/getRichTextPublic [get]
func (a *richText) FindRichTextByName(c *gin.Context) {
	name := c.Query("name")
	richText, err := serviceRichText.FindRichTextByName(name)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(richText, c)
}
