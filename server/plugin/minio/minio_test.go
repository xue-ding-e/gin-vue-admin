package minio

//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/minio/minio-go/v7/pkg/credentials"
//)
//
//const (
//	endpoint        = "http://203.56.201.15:65090/"      // MinIO 服务地址
//	accessKeyID     = "gva"                              // MinIO 管理员 Access Key
//	secretAccessKey = "#cilG+XzToR{*=~8@B7PBI]t&{Q`C?hU" // MinIO 管理员 Secret Key
//)
//
//// GetCredentialsHandler 处理获取临时凭证的请求
//func GetCredentialsHandler(w http.ResponseWriter, r *http.Request) {
//	// 获取上传对象名称，从请求中获取或由后端指定
//	objectName := "uploads/user123/upload.jpg"
//
//	// 生成策略，限制只能上传指定的对象
//	policy := fmt.Sprintf(`{
// "Version": "2012-10-17",
// "Statement": [
//   {
//     "Effect": "Allow",
//     "Action": [
//       "s3:GetObject",
//       "s3:PutObject",
//       "s3:DeleteObject",
//       "s3:GetBucketLocation"
//     ],
//     "Resource": [
//       "arn:aws:s3:::%s"
//     ]
//   }
// ]
//}`, objectName)
//
//	// 获取临时凭证
//	creds, err := getTemporaryCredentials(policy)
//	if err != nil {
//		http.Error(w, "获取临时凭证失败", http.StatusInternalServerError)
//		log.Printf("获取临时凭证失败：%v", err)
//		return
//	}
//
//	// 将临时凭证和对象名称一起发送给前端
//	response := map[string]interface{}{
//		"credentials": creds,
//		"objectName":  objectName,
//		"bucketName":  "test",
//		"endpoint":    "http://203.56.201.15:65090", // 前端访问的 MinIO 域名
//		"useSSL":      false,                        // 是否使用 HTTPS
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(response)
//}
//
//// 获取临时凭证
//func getTemporaryCredentials(policy string) (map[string]string, error) {
//	stsEndpoint := endpoint
//
//	// 创建 MinIO 客户端，使用管理员凭证
//	stsCreds := credentials.NewStaticV4(accessKeyID, secretAccessKey, "")
//	stsOpts := credentials.STSAssumeRoleOptions{
//		Policy:          policy,
//		DurationSeconds: 3600, // 有效期 1 小时
//	}
//
//	stsProvider, err := credentials.NewSTSAssumeRole(stsEndpoint, stsOpts)
//	if err != nil {
//		return nil, err
//	}
//
//	creds, err := stsProvider.Retrieve()
//	if err != nil {
//		return nil, err
//	}
//
//	credentials := map[string]string{
//		"accessKey":    creds.AccessKeyID,
//		"secretKey":    creds.SecretAccessKey,
//		"sessionToken": creds.SessionToken,
//	}
//
//	return credentials, nil
//}
