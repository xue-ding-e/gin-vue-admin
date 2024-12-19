package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Good = new(good)

type good struct{}

// CreateGood 创建商品
// @Tags Good
// @Summary 创建商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Good true "创建商品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /good/createGood [post]
func (a *good) CreateGood(c *gin.Context) {
	var info shop.Good
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodService.CreateGood(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteGood 删除商品
// @Tags Good
// @Summary 删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Good true "删除商品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /good/deleteGood [delete]
func (a *good) DeleteGood(c *gin.Context) {
	ID := c.Query("ID")
	err := goodService.DeleteGood(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGoodByIds 批量删除商品
// @Tags Good
// @Summary 批量删除商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /good/deleteGoodByIds [delete]
func (a *good) DeleteGoodByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := goodService.DeleteGoodByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGood 更新商品
// @Tags Good
// @Summary 更新商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Good true "更新商品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /good/updateGood [put]
func (a *good) UpdateGood(c *gin.Context) {
	var info shop.Good
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodService.UpdateGood(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGood 用id查询商品
// @Tags Good
// @Summary 用id查询商品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Good true "用id查询商品"
// @Success 200 {object} response.Response{data=model.Good,msg=string} "查询成功"
// @Router /good/findGood [get]
func (a *good) FindGood(c *gin.Context) {
	ID := c.Query("ID")
	regood, err := goodService.GetGood(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(regood, c)
}

// GetGoodList 分页获取商品列表
// @Tags Good
// @Summary 分页获取商品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GoodSearch true "分页获取商品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /good/getGoodList [get]
func (a *good) GetGoodList(c *gin.Context) {
	var pageInfo request.GoodSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := goodService.GetGoodInfoList(pageInfo)
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

// GetGoodPublic 不需要鉴权的商品接口
// @Tags Good
// @Summary 不需要鉴权的商品接口
// @accept application/json
// @Produce application/json
// @Param data query request.GoodSearch true "分页获取商品列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /good/getGoodPublic [get]
func (a *good) GetGoodPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	goodService.GetGoodPublic()
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的商品接口信息"}, "获取成功", c)
}
