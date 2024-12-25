package api

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Minio = new(MinioApi)

type MinioApi struct{}

func (self *MinioApi) GetMinioSts(c *gin.Context) {
	ids := []string{"1", "2", "3", "4", "5", "6"}
	var resources []string
	for _, id := range ids {
		resource := fmt.Sprintf("arn:aws:s3:::test/%s/*", id)
		resources = append(resources, resource)
	}

	policyObj := gabs.New()
	policyObj.Set("2012-10-17", "Version")

	statement := map[string]interface{}{
		"Effect": "Allow",
		"Action": []string{
			"s3:GetObject",
			"s3:GetBucketLocation",
		},
		"Resource": resources,
	}

	policyObj.Array("Statement")
	policyObj.ArrayAppend(statement, "Statement")

	policy := policyObj.String()

	credentials, err := serviceMinio.GetMinioSts(policy)
	if err != nil {
		global.GVA_LOG.Error("获取 MinIO STS 凭证失败", zap.Error(err))
		response.FailWithMessage("获取 MinIO STS 凭证失败", c)
		return
	}
	response.OkWithData(credentials, c)
}
