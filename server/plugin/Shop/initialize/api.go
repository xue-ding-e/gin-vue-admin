package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{ApiGroup: "商品分类", Method: "POST", Path: "/category/createCategory", Description: "新增商品分类"},
		{ApiGroup: "商品分类", Method: "DELETE", Path: "/category/deleteCategory", Description: "删除商品分类"},
		{ApiGroup: "商品分类", Method: "DELETE", Path: "/category/deleteCategoryByIds", Description: "批量删除商品分类"},
		{ApiGroup: "商品分类", Method: "PUT", Path: "/category/updateCategory", Description: "更新商品分类"},
		{ApiGroup: "商品分类", Method: "GET", Path: "/category/findCategory", Description: "根据ID获取商品分类"},
		{ApiGroup: "商品分类", Method: "GET", Path: "/category/getCategoryList", Description: "获取商品分类列表"},

		{ApiGroup: "商品", Method: "POST", Path: "/good/createGood", Description: "新增商品"},
		{ApiGroup: "商品", Method: "DELETE", Path: "/good/deleteGood", Description: "删除商品"},
		{ApiGroup: "商品", Method: "DELETE", Path: "/good/deleteGoodByIds", Description: "批量删除商品"},
		{ApiGroup: "商品", Method: "PUT", Path: "/good/updateGood", Description: "更新商品"},
		{ApiGroup: "商品", Method: "GET", Path: "/good/findGood", Description: "根据ID获取商品"},
		{ApiGroup: "商品", Method: "GET", Path: "/good/getGoodList", Description: "获取商品列表"},

		{ApiGroup: "sku", Method: "POST", Path: "/sku/createSku", Description: "新增sku"},
		{ApiGroup: "sku", Method: "DELETE", Path: "/sku/deleteSku", Description: "删除sku"},
		{ApiGroup: "sku", Method: "DELETE", Path: "/sku/deleteSkuByIds", Description: "批量删除sku"},
		{ApiGroup: "sku", Method: "PUT", Path: "/sku/updateSku", Description: "更新sku"},
		{ApiGroup: "sku", Method: "GET", Path: "/sku/findSku", Description: "根据ID获取sku"},
		{ApiGroup: "sku", Method: "GET", Path: "/sku/getSkuList", Description: "获取sku列表"},

		{ApiGroup: "购物车", Method: "POST", Path: "/cart/createCart", Description: "新增购物车"},
		{ApiGroup: "购物车", Method: "DELETE", Path: "/cart/deleteCart", Description: "删除购物车"},
		{ApiGroup: "购物车", Method: "DELETE", Path: "/cart/deleteCartByIds", Description: "批量删除购物车"},
		{ApiGroup: "购物车", Method: "PUT", Path: "/cart/updateCart", Description: "更新购物车"},
		{ApiGroup: "购物车", Method: "GET", Path: "/cart/findCart", Description: "根据ID获取购物车"},
		{ApiGroup: "购物车", Method: "GET", Path: "/cart/getCartList", Description: "获取购物车列表"},
		{ApiGroup: "购物车", Method: "POST", Path: "/cart/cutCart", Description: "删除购物车"},
		{ApiGroup: "购物车", Method: "POST", Path: "/cart/addCart", Description: "添加购物车"},
		{ApiGroup: "购物车", Method: "GET", Path: "/cart/getSelfCart", Description: "获取自身购物车"},
		{ApiGroup: "购物车", Method: "GET", Path: "/cart/clearCart", Description: "全部删除购物车"},
	}
	utils.RegisterApis(entities, true)
}
