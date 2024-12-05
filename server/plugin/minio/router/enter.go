package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/minio/api"

var (
	Router   = new(router)
	apiMinio = api.Api.Minio
)

type router struct{ Minio MI }
