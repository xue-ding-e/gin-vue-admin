package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name string `json:"name"`
}

func (t *TestStruct) Validate(c *gin.Context) error {
	if t.Name == "invalid" {
		return errors.New("invalid name")
	}
	return nil
}

// 普通结构体，没有实现 ValidatorHook
type NormalStruct struct {
	Age int `json:"age"`
}

func TestShouldBindWith(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("with validator hook", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 创建测试请求
		body := strings.NewReader(`{"name": "invalid"}`)
		c.Request, _ = http.NewRequest("POST", "/", body)

		obj := &TestStruct{}
		err := c.ShouldBindWith(obj, binding.JSON)
		assert.Error(t, err)
		assert.Equal(t, "invalid name", err.Error())
	})

	t.Run("normal struct binding", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := strings.NewReader(`{"age": 18}`)
		c.Request, _ = http.NewRequest("POST", "/", body)

		obj := &NormalStruct{}
		err := c.ShouldBindWith(obj, binding.JSON)
		assert.NoError(t, err)
		assert.Equal(t, 18, obj.Age)
	})
}
