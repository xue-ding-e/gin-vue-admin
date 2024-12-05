package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/minio/service"

var (
	Api          = new(api)
	serviceMinio = service.Service.Minio
)

type api struct{ Minio MinioApi }
