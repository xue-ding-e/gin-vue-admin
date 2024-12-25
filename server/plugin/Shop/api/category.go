package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	shop "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Category = new(category)

type category struct{}

// CreateCategory 创建商品分类
// @Tags Category
// @Summary 创建商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Category true "创建商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /category/createCategory [post]
func (self *category) CreateCategory(c *gin.Context) {
	var category shop.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := categoryService.CreateCategory(&category); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCategory 删除商品分类
// @Tags Category
// @Summary 删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Category true "删除商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /category/deleteCategory [delete]
func (self *category) DeleteCategory(c *gin.Context) {
	ID := c.Query("ID")
	if err := categoryService.DeleteCategory(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCategoryByIds 批量删除商品分类
// @Tags Category
// @Summary 批量删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /category/deleteCategoryByIds [delete]
func (self *category) DeleteCategoryByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := categoryService.DeleteCategoryByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCategory 更新商品分类
// @Tags Category
// @Summary 更新商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.Category true "更新商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /category/updateCategory [put]
func (self *category) UpdateCategory(c *gin.Context) {
	var category shop.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := categoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCategory 用id查询商品分类
// @Tags Category
// @Summary 用id查询商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.Category true "用id查询商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /category/findCategory [get]
func (self *category) FindCategory(c *gin.Context) {
	ID := c.Query("ID")
	if recategory, err := categoryService.GetCategory(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recategory": recategory}, c)
	}
}

// GetCategoryList 分页获取商品分类列表
// @Tags Category
// @Summary 分页获取商品分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.CategorySearch true "分页获取商品分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
func (self *category) GetCategoryList(c *gin.Context) {
	var pageInfo request.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := categoryService.GetCategoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (self *category) GetCategoryMobile(c *gin.Context) {
	if data, err := categoryService.GetCategoryMobile(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(data, "获取成功", c)
	}
}

// GetCategoryPublic 不需要鉴权的商品分类接口
// @Tags Category
// @Summary 不需要鉴权的商品分类接口
// @accept application/json
// @Produce application/json
// @Param data query shopReq.CategorySearch true "分页获取商品分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /category/getCategoryList [get]
func (self *category) GetCategoryPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品分类接口信息",
	}, "获取成功", c)
}
