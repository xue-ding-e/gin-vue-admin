package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body []byte
		var userId int

		// 如果不是 GET 请求，则读取 Body 并重置，便于后续逻辑或其他中间件可再次读取
		if c.Method() != http.MethodGet {
			body = c.Body()
			_ = c.Request().SetBody(body)
		} else {
			// 这里模拟在原 Gin 中获取并序列化查询参数的做法
			query := c.Context().QueryArgs().String()
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
			_ = c.Request().SetBody(body)
		}

		// 读取用户 ID
		claims, _ := utils.GetClaims(c) // 需修改 utils.GetClaims 以兼容 Fiber
		if claims != nil && claims.BaseClaims.ID != 0 {
			userId = int(claims.BaseClaims.ID)
		} else {
			id, err := strconv.Atoi(c.Get("x-user-id"))
			if err != nil {
				userId = 0
			} else {
				userId = id
			}
		}

		// 构造操作记录
		record := system.SysOperationRecord{
			Ip:     c.IP(),
			Method: c.Method(),
			Path:   c.Path(),
			Agent:  c.Get("User-Agent"),
			Body:   "",
			UserID: userId,
		}

		// 判断是否上传文件
		if strings.Contains(c.Get("Content-Type"), "multipart/form-data") {
			record.Body = "[文件]"
		} else {
			if len(body) > bufferSize {
				record.Body = "[超出记录长度]"
			} else {
				record.Body = string(body)
			}
		}

		// 开始计时，执行下一个处理
		startTime := time.Now()
		if err := c.Next(); err != nil {
			// Fiber 默认错误处理可在 app.Use(recover.New()) 或自定义错误处理中
			// 这里可自行决定如何记录错误信息
			global.GVA_LOG.Error("fiber next error:", zap.Error(err))
		}
		latency := time.Since(startTime)

		// 记录响应
		// 在 Fiber 中可直接通过 c.Response().Body() 获取响应内容
		record.Status = c.Response().StatusCode()
		record.ErrorMessage = "" // Fiber 并没有像 Gin 那样的 c.Errors，可自定义错误记录机制
		record.Latency = latency
		record.Resp = string(c.Response().Body())

		// 检查特定响应头
		respHeaders := []string{
			c.GetRespHeader("Pragma"),
			c.GetRespHeader("Expires"),
			c.GetRespHeader("Cache-Control"),
			c.GetRespHeader("Content-Type"),
			c.GetRespHeader("Content-Disposition"),
			c.GetRespHeader("Content-Transfer-Encoding"),
		}
		for _, h := range respHeaders {
			if strings.Contains(h, "public") ||
				strings.Contains(h, "0") ||
				strings.Contains(h, "must-revalidate") ||
				strings.Contains(h, "application/force-download") ||
				strings.Contains(h, "application/octet-stream") ||
				strings.Contains(h, "application/vnd.ms-excel") ||
				strings.Contains(h, "application/download") ||
				strings.Contains(h, "attachment") ||
				strings.Contains(h, "binary") {
				if len(record.Resp) > bufferSize {
					record.Body = "超出记录长度"
				}
				break
			}
		}

		// 写入操作记录
		if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Error(err))
		}

		return nil
	}
}
