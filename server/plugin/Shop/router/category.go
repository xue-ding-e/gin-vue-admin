package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Category = new(category)

type category struct{}

// Init 初始化 liCategory表 路由信息
func (self *category) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("category").Use(middleware.OperationRecord())
		group.PUT("updateCategory", apiCategory.UpdateCategory)
		group.POST("createCategory", apiCategory.CreateCategory)
		group.POST("deleteCategory", apiCategory.DeleteCategory) // 使用query参数 id
	}
	{
		group := private.Group("category")
		group.GET("getAllChildren", apiCategory.GetAllChildren) // 使用query参数 parentId
		group.GET("getAllCategories", apiCategory.GetAllCategories)
		group.POST("disableCategory", apiCategory.DisableCategory)  // 使用请求体参数
		group.GET("findCategoryById", apiCategory.FindCategoryById) // 使用query参数 id
	}
	{
		//group := public.Group("category")
		//group.GET("getAllCategories", apiCategory.GetAllCategories) // liCategory表开放接口
	}
}
