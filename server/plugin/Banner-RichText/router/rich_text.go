package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/xue-ding-e/gin"
)

var RichText = new(richText)

type richText struct{}

// Init 初始化 富文本 路由信息
func (r *richText) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("richText").Use(middleware.OperationRecord())
		group.POST("createRichText", apiRichText.CreateRichText)             // 新建富文本
		group.DELETE("deleteRichText", apiRichText.DeleteRichText)           // 删除富文本
		group.DELETE("deleteRichTextByIds", apiRichText.DeleteRichTextByIds) // 批量删除富文本
		group.PUT("updateRichText", apiRichText.UpdateRichText)              // 更新富文本
	}
	{
		group := private.Group("richText")
		group.GET("findRichText", apiRichText.FindRichText)       // 根据ID获取富文本
		group.GET("getRichTextList", apiRichText.GetRichTextList) // 获取富文本列表
	}
	{
		group := public.Group("richText")
		group.GET("findRichTextByName", apiRichText.FindRichTextByName) // 富文本开放接口
	}
}
