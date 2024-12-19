package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Address/api"

var (
	Router     = new(router)
	addressApi = api.Api.Address
)

type router struct{ Address address }
