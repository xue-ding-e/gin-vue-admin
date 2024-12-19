package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/api"

var (
	Router      = new(router)
	apiBanner   = api.Api.Banner
	apiRichText = api.Api.RichText
)

type router struct {
	Banner   banner
	RichText richText
}
