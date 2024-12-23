package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"reflect"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SYSTEM_ERROR = 110
	ERROR        = 7
	SUCCESS      = 0
)

func Result(code int, data interface{}, msg string, c *fiber.Ctx) {
	// 开始时间
	c.Status(http.StatusOK).JSON(Response{
		code,
		data,
		msg,
	})
}

func Ok(c *fiber.Ctx) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *fiber.Ctx) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *fiber.Ctx) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailSystemError(c *fiber.Ctx) {
	Result(SYSTEM_ERROR, map[string]interface{}{}, "系统错误", c)
}

func FailWithMessage(message string, c *fiber.Ctx) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailSystemErrorWithMessage(message string, c *fiber.Ctx) {
	Result(SYSTEM_ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *fiber.Ctx) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(ERROR, data, message, c)
}

func FailSystemErrorWithDetailed(data interface{}, message string, c *fiber.Ctx) {
	Result(SYSTEM_ERROR, data, message, c)
}

type ResponseV2 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ResultV2(code int, data interface{}, msg string, c *fiber.Ctx) {
	responseData := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	if data != nil {
		val := reflect.ValueOf(data)
		typ := reflect.TypeOf(data)

		switch val.Kind() {
		case reflect.Struct:
			for i := 0; i < val.NumField(); i++ {
				fieldName := typ.Field(i).Tag.Get("json")
				if fieldName == "" {
					fieldName = typ.Field(i).Name
				}
				responseData[fieldName] = val.Field(i).Interface()
			}
		case reflect.Map:
			for _, key := range val.MapKeys() {
				responseData[key.String()] = val.MapIndex(key).Interface()
			}
		default:
			responseData["data"] = data
		}
	}
	c.JSON(http.StatusOK, responseData)
}

func OkV2(c *fiber.Ctx) {
	ResultV2(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessageV2(message string, c *fiber.Ctx) {
	ResultV2(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithDataV2(data interface{}, c *fiber.Ctx) {
	ResultV2(SUCCESS, data, "查询成功", c)
}

func OkWithDetailedV2(data interface{}, message string, c *fiber.Ctx) {
	ResultV2(SUCCESS, data, message, c)
}

func FailV2(c *fiber.Ctx) {
	ResultV2(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessageV2(message string, c *fiber.Ctx) {
	ResultV2(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailedV2(data interface{}, message string, c *fiber.Ctx) {
	ResultV2(ERROR, data, message, c)
}
