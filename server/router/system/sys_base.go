package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/utils"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", utils.GvaBaseLoginApi)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
