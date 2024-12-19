package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{{Path: "/banner/createBanner", Description: "新增轮播图", ApiGroup: "轮播图", Method: "POST"},
		{Path: "/banner/deleteBanner", Description: "删除轮播图", ApiGroup: "轮播图", Method: "DELETE"},
		{Path: "/banner/deleteBannerByIds", Description: "批量删除轮播图", ApiGroup: "轮播图", Method: "DELETE"},
		{Path: "/banner/updateBanner", Description: "更新轮播图", ApiGroup: "轮播图", Method: "PUT"},
		{Path: "/banner/findBanner", Description: "根据ID获取轮播图", ApiGroup: "轮播图", Method: "GET"},
		{Path: "/banner/getBannerList", Description: "获取轮播图列表", ApiGroup: "轮播图", Method: "GET"},
		{Path: "/banner/getBannerPublic", Description: "获取广告", ApiGroup: "轮播图", Method: "GET"},
		{Path: "/richText/createRichText", Description: "新增富文本", ApiGroup: "富文本", Method: "POST"},
		{Path: "/richText/deleteRichText", Description: "删除富文本", ApiGroup: "富文本", Method: "DELETE"},
		{Path: "/richText/deleteRichTextByIds", Description: "批量删除富文本", ApiGroup: "富文本", Method: "DELETE"},
		{Path: "/richText/updateRichText", Description: "更新富文本", ApiGroup: "富文本", Method: "PUT"},
		{Path: "/richText/findRichText", Description: "根据ID获取富文本", ApiGroup: "富文本", Method: "GET"},
		{Path: "/richText/getRichTextList", Description: "获取富文本列表", ApiGroup: "富文本", Method: "GET"},
		{Path: "/richText/findRichTextByName", Description: "c端用户根据name查找富文本", ApiGroup: "富文本", Method: "GET"}}
	utils.RegisterApis(entities, true)
}
