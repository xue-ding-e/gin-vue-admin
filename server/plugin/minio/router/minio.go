package router

import (
	"github.com/gin-gonic/gin"
)

var Minio = new(MI)

type MI struct{}

// Init 初始化 minio存储 路由信息
func (r *MI) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		//_ := private.Group("MI").Use(middleware.OperationRecord())
	}
	{
		group := private.Group("minio")
		group.GET("getMinioSts", apiMinio.GetMinioSts) // 获取minio sts凭证
		//group.GET("getPresignedUrl", apiMinio.GetPresignedUrl)                 // minio获取上传预签名 URL
		//group.GET("getDownloadPresignedUrl", apiMinio.GetDownloadPresignedUrl) // minio获取下载预签名 URL

	}
	{
		// 测试用
		// group := public.Group("minio")
		// group.GET("getMinioSts", apiMinio.GetMinioSts) // 获取minio sts凭证
	}
}
