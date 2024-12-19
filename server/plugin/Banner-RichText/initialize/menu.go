package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{{ParentId: 0, Path: "BannerMenu", Name: "BannerMenu", Hidden: false, Component: "view/routerHolder.vue", Sort: 0, Meta: model.Meta{Title: "轮播图和富文本", Icon: "school"}},
		{ParentId: 0, Path: "banner", Name: "banner", Hidden: false, Component: "plugin/Banner-RichText/view/banner.vue", Sort: 1, Meta: model.Meta{Title: "轮播图", Icon: "picture-filled"}},
		{ParentId: 0, Path: "richText", Name: "richText", Hidden: false, Component: "plugin/Banner-RichText/view/richText.vue", Sort: 0, Meta: model.Meta{Title: "富文本", Icon: ""}}}
	utils.RegisterMenus(entities, true)
}
