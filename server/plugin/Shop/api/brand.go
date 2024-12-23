package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Brand = new(brand)

type brand struct{}

// CreateBrand 创建品牌表
// @Tags Brand
// @Summary 创建品牌表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "创建品牌表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /brand/createBrand [post]
func (self *brand) CreateBrand(c *fiber.Ctx) {
	var info model.Brand
	err := c.BodyParser(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.CreatedBy = utils.GetUserID(c)
	err = serviceBrand.CreateBrand(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBrand 删除品牌表
// @Tags Brand
// @Summary 删除品牌表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "删除品牌表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /brand/deleteBrand [delete]
func (self *brand) DeleteBrand(c *fiber.Ctx) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := serviceBrand.DeleteBrand(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBrandByIds 批量删除品牌表
// @Tags Brand
// @Summary 批量删除品牌表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /brand/deleteBrandByIds [delete]
func (self *brand) DeleteBrandByIds(c *fiber.Ctx) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := serviceBrand.DeleteBrandByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBrand 更新品牌表
// @Tags Brand
// @Summary 更新品牌表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "更新品牌表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /brand/updateBrand [put]
func (self *brand) UpdateBrand(c *fiber.Ctx) {
	var info model.Brand
	err := c.BodyParser(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.UpdatedBy = utils.GetUserID(c)
	err = serviceBrand.UpdateBrand(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBrand 用id查询品牌表
// @Tags Brand
// @Summary 用id查询品牌表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Brand true "用id查询品牌表"
// @Success 200 {object} response.Response{data=model.Brand,msg=string} "查询成功"
// @Router /brand/findBrand [get]
func (self *brand) FindBrand(c *fiber.Ctx) {
	ID := c.Query("ID")
	rebrand, err := serviceBrand.GetBrand(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebrand, c)
}

// GetBrandList 分页获取品牌表列表
// @Tags Brand
// @Summary 分页获取品牌表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.BrandSearch true "分页获取品牌表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /brand/getBrandList [get]
func (self *brand) GetBrandList(c *fiber.Ctx) {
	var pageInfo request.BrandSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceBrand.GetBrandInfoList(pageInfo)
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

// GetBrandPublic 不需要鉴权的品牌表接口
// @Tags Brand
// @Summary 不需要鉴权的品牌表接口
// @accept application/json
// @Produce application/json
// @Param data query request.BrandSearch true "分页获取品牌表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /brand/getBrandPublic [get]
func (self *brand) GetBrandPublic(c *fiber.Ctx) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceBrand.GetBrandPublic()
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的品牌表接口信息"}, "获取成功", c)
}
