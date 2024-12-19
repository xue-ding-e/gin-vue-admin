package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/service"

var (
	Api             = new(api)
	goodService     = service.Service.Good
	categoryService = service.Service.Category
)

type api struct {
	Good     good
	Category category
}
