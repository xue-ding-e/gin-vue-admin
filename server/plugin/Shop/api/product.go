package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
	"go.uber.org/zap"
)

var Product = new(product)

type product struct{}

// CreateProduct 创建产品
// @Tags Product
// @Summary 创建产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建产品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /product/createProduct [post]
func (self *product) CreateProduct(c *fiber.Ctx) {
	var info model.Product
	err := c.BodyParser(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceProduct.CreateProduct(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProduct 删除产品
// @Tags Product
// @Summary 删除产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除产品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /product/deleteProduct [delete]
func (self *product) DeleteProduct(c *fiber.Ctx) {
	ID := c.Query("ID")
	err := serviceProduct.DeleteProduct(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductByIds 批量删除产品
// @Tags Product
// @Summary 批量删除产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /product/deleteProductByIds [delete]
func (self *product) DeleteProductByIds(c *fiber.Ctx) {
	IDs := c.QueryArray("ID")
	err := serviceProduct.DeleteProductByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProduct 更新产品
// @Tags Product
// @Summary 更新产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新产品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /product/updateProduct [put]
func (self *product) UpdateProduct(c *fiber.Ctx) {
	var info model.Product
	err := c.BodyParser(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceProduct.UpdateProduct(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProduct 用id查询产品
// @Tags Product
// @Summary 用id查询产品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Product true "用id查询产品"
// @Success 200 {object} response.Response{data=model.Product,msg=string} "查询成功"
// @Router /product/findProduct [get]
func (self *product) FindProduct(c *fiber.Ctx) {
	ID := c.Query("ID")
	reproduct, err := serviceProduct.GetProduct(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reproduct, c)
}

// GetProductList 分页获取产品列表
// @Tags Product
// @Summary 分页获取产品列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.ProductSearch true "分页获取产品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /product/getProductList [get]
func (self *product) GetProductList(c *fiber.Ctx) {
	var pageInfo request.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceProduct.GetProductInfoList(pageInfo)
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

// 根据分类 ID 获取商品列表
func (self *product) GetProductsByCategoryID(c *fiber.Ctx) {
	categoryIDStr := c.Query("categoryId")
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的 categoryId", c)
		return
	}
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	products, err := serviceProduct.GetProductsByCategoryID(categoryID, page, size)
	if err != nil {
		response.FailWithMessage("获取商品列表失败", c)
		return
	}
	response.OkWithData(products, c)
}

// 根据商品 ID 获取商品详情
func (self *product) GetProductInfo(c *fiber.Ctx) {
	prodIDStr := c.Query("prodId")
	prodID, err := strconv.ParseUint(prodIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的 prodId", c)
		return
	}

	product_data, err := serviceProduct.GetProductByID(prodID)
	if err != nil {
		response.FailWithMessage("商品不存在", c)
		return
	}

	// 获取 SKU 列表
	skuList, err := serviceProduct.GetSkuListByProdID(prodID)
	if err != nil {
		response.FailWithMessage("获取 SKU 列表失败", c)
		return
	}
	product_data.SkuList = skuList

	// 处理配送方式和运费模板（此处省略，需要根据实际情况实现）

	response.OkWithData(product_data, c)
}
