package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/api"

var (
	Router  = new(router)
	apiGood = api.Api.Good
)

type router struct{ Good good }
