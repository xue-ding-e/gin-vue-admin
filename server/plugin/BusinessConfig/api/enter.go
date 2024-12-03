package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/BusinessConfig/service"

var (
	Api                   = new(api)
	serviceBusinessConfig = service.Service.BusinessConfig
)

type api struct{ BusinessConfig businessConfig }
