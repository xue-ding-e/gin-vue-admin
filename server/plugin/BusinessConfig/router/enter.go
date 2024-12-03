package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/api"

var (
	Router            = new(router)
	apiBusinessConfig = api.Api.BusinessConfig
)

type router struct{ BusinessConfig businessConfig }
