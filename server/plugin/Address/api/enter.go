package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/Address/service"

var (
	Api            = new(api)
	addressService = service.Service.Address
)

type api struct{ Address AddressApi }
