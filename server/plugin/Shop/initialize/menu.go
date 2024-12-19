package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "shop", Name: "shop", Component: "view/routerHolder.vue", Sort: 1000, Meta: model.Meta{Title: "商城管理", Icon: "shop"}},
		{MenuLevel: 0, Hidden: false, ParentId: 36, Path: "category", Name: "category", Component: "plugin/Shop/view/category.vue", Sort: 2000, Meta: model.Meta{Title: "商品分类", Icon: "cherry"}},
		{MenuLevel: 0, Hidden: false, ParentId: 36, Path: "good", Name: "good", Component: "plugin/Shop/view/good.vue", Sort: 1000, Meta: model.Meta{Title: "商品管理", Icon: "shopping-bag"}},
		{MenuLevel: 0, Hidden: true, ParentId: 36, Path: "sku", Name: "sku", Component: "plugin/Shop/view/sku.vue", Sort: 0, Meta: model.Meta{Title: "sku-${id}", Icon: ""}},
		{MenuLevel: 0, Hidden: false, ParentId: 36, Path: "order", Name: "order", Component: "plugin/Shop/view/order.vue", Sort: 0, Meta: model.Meta{Title: "订单管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: true, ParentId: 0, Path: "cart", Name: "cart", Component: "plugin/Shop/view/cart.vue", Sort: 0, Meta: model.Meta{Title: "购物车", Icon: ""}},
	}
	utils.RegisterMenus(entities, true)
}
