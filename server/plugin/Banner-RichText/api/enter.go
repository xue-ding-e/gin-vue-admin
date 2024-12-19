package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/service"

var (
	Api             = new(api)
	bannerService   = service.Service.Banner
	serviceRichText = service.Service.RichText
)

type api struct {
	Banner   banner
	RichText richText
}
