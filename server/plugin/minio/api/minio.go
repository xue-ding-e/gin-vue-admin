package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Minio = new(MinioApi)

type MinioApi struct{}

func (self *MinioApi) GetMinioSts(c *gin.Context) {
	credentials, err := serviceMinio.GetMinioSts("")
	if err != nil {
		global.GVA_LOG.Error("获取minio sts凭证失败", zap.Error(err))
		response.FailWithMessage("获取minio sts凭证失败", c)
		return
	}
	response.OkWithData(credentials, c)
}
