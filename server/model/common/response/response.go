package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type ResponseNoMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

const (
	SYSTEM_ERROR = 110
	ERROR        = 7
	SUCCESS      = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func ResultNoMessage(code int, data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, ResponseNoMessage{
		code,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func Ok2(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailSystemError(c *gin.Context) {
	Result(SYSTEM_ERROR, map[string]interface{}{}, "系统错误", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailSystemErrorWithMessage(message string, c *gin.Context) {
	Result(SYSTEM_ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailSystemErrorWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SYSTEM_ERROR, data, message, c)
}
