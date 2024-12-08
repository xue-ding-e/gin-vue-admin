package service

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/minio/plugin"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var Minio = new(MI)

type MI struct{}

// PRODUCTION 自己根据业务改造
func (a *MI) GetMinioSts(policy string) (map[string]string, error) {
	if policy == "" {
		// 生成策略，限制上传指定的对象  注:一定要以\的形式 不能直接以多行字符串的形式否则会报错
		//TODO policy封装动态配置
		policy = `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::test/*"
            ]
        }
    ]
}`
	}
	// 获取临时凭证
	creds, err := getTemporaryCredentials(policy)
	if err != nil {
		global.GVA_LOG.Debug("获取临时凭证失败", zap.Error(err))
		return nil, err
	}
	return creds, nil
}

func getTemporaryCredentials(policy string) (map[string]string, error) {

	// 创建 STSAssumeRoleOptions，并在其中包含管理员凭证
	stsOpts := credentials.STSAssumeRoleOptions{
		AccessKey:       plugin.Minio.Username,        // 管理员 账号
		SecretKey:       plugin.Minio.Password,        // 管理员 密码
		DurationSeconds: plugin.Minio.DurationSeconds, // 有效期，单位为秒 , 最大为43200(12小时)
		Policy:          policy,                       // 策略
		Location:        plugin.Minio.Region,
		//RoleARN:         "",
		RoleSessionName: "anysession",
	}

	// 创建 STSAssumeRole 对象 , 注意生成对象失败err依然是nil所以获取凭证失败可能是
	// 创建对象失败 , 并不是下面获取临时凭证.Get值的时候失败 , 这个时候请检查上面的
	// STSAssumeRoleOptions 配置信息
	credsProvider, err := credentials.NewSTSAssumeRole(plugin.Minio.Endpoint, stsOpts)
	if err != nil {
		return nil, fmt.Errorf("创建 STSAssumeRole 失败: %v", err)
	}

	// 获取临时凭证
	credsValue, err := credsProvider.Get()
	if err != nil {
		return nil, fmt.Errorf("获取临时凭证失败: %v", err)
	}

	credentials := map[string]string{
		"accessKey":    credsValue.AccessKeyID,
		"secretKey":    credsValue.SecretAccessKey,
		"sessionToken": credsValue.SessionToken,
	}
	return credentials, nil
}

// 预签名
//func (a *MI) GetPresignedUrl(c *gin.Context) {
//	// 获取上传对象名称，可以从请求参数中获取
//	filename := c.Query("filename")
//	if filename == "" {
//		response.FailWithMessage("filename 不能为空", c)
//		return
//	}
//	userid := utils.GetUserID(c)
//	if userid == 0 {
//		response.FailWithMessage("userid 不能为空", c)
//		return
//	}
//	filename = "uploads/" + utils2.UintToString(userid) + "/" + filename // 默认对象名称
//
//	// 初始化 MinIO 客户端
//	minioClient, err := minio.New(global.GVA_CONFIG.Minio.Endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(global.GVA_CONFIG.Minio.AccessKeyId, global.GVA_CONFIG.Minio.AccessKeySecret, ""),
//		Secure: global.GVA_CONFIG.Minio.UseSSL,
//	})
//	if err != nil {
//		global.GVA_LOG.Debug("创建 MinIO 客户端失败", zap.Error(err))
//		response.FailWithMessage("创建 MinIO 客户端失败", c)
//		return
//	}
//
//	// 检查存储桶是否存在
//	exists, err := minioClient.BucketExists(c, global.GVA_CONFIG.Minio.BucketName2)
//	if err != nil {
//		global.GVA_LOG.Debug("检查存储桶失败", zap.Error(err))
//		response.FailWithMessage("检查存储桶失败", c)
//		return
//	}
//	if !exists {
//		// 如果存储桶不存在，则创建
//		err = minioClient.MakeBucket(c, global.GVA_CONFIG.Minio.BucketName2, minio.MakeBucketOptions{Region: "cn-north-1"})
//		if err != nil {
//			global.GVA_LOG.Debug("创建存储桶失败", zap.Error(err))
//			response.FailWithMessage("创建存储桶失败", c)
//			return
//		}
//	}
//
//	// 生成预签名的上传 URL，有效期为1小时
//	presignedURL, err := minioClient.PresignedPutObject(c, global.GVA_CONFIG.Minio.BucketName2, filename, time.Hour)
//	if err != nil {
//		global.GVA_LOG.Debug("生成预签名 URL 失败", zap.Error(err))
//		response.FailWithMessage("生成预签名 URL 失败", c)
//		return
//	}
//
//	response.OkWithData(gin.H{
//		"presignedUrl": presignedURL.String(),
//		"objectName":   filename,
//		"bucketName":   global.GVA_CONFIG.Minio.BucketName2,
//	}, c)
//}
//
//// TODO 文件名字逻辑做后端生成
//// 获取查看（下载）预签名 URL（有效期 1 小时）
//func (a *MI) GetDownloadPresignedUrl(c *gin.Context) {
//	// 获取参数
//	objectName := c.Query("filename")
//	if objectName == "" {
//		objectName = "uploads/user123/upload.jpg" // 默认对象名称
//	}
//	bucketName := "test"
//
//	// 初始化 MinIO 客户端
//	minioClient, err := minio.New(global.GVA_CONFIG.Minio.Endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(global.GVA_CONFIG.Minio.AccessKeyId, global.GVA_CONFIG.Minio.AccessKeySecret, ""),
//		Secure: global.GVA_CONFIG.Minio.UseSSL,
//		Region: "cn-north-1",
//	})
//	if err != nil {
//		global.GVA_LOG.Debug("创建 MinIO 客户端失败", zap.Error(err))
//		response.FailWithMessage("创建 MinIO 客户端失败", c)
//		return
//	}
//
//	// 生成下载预签名 URL，有效期为 1 小时
//	presignedURL, err := minioClient.PresignedGetObject(c, bucketName, objectName, time.Hour, nil)
//	if err != nil {
//		global.GVA_LOG.Debug("生成下载预签名 URL 失败", zap.Error(err))
//		response.FailWithMessage("生成下载预签名 URL 失败", c)
//		return
//	}
//
//	response.OkWithData(gin.H{
//		"presignedUrl": presignedURL.String(),
//		"objectName":   objectName,
//		"bucketName":   bucketName,
//	}, c)
//}
