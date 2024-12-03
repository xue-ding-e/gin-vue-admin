package config

import (
	"errors"
)

var GLOBAL_MAP map[string]interface{}

var GLOBAL_CONFIG GLOBAL_CONFIG_TYPE

type GLOBAL_CONFIG_TYPE struct {
	G文件上传最大值 *int `json:"文件上传最大值" form:"文件上传最大值" mapstructure:"文件上传最大值"` // 文件上传最大值(MB)
}

// 校验函数
func (self *GLOBAL_CONFIG_TYPE) ValidateBusinessConfig() {
	if self.G文件上传最大值 == nil || *self.G文件上传最大值 <= 0 {
		panic("文件上传最大值必须大于 0")
	}
}

// 校验函数
func (self *GLOBAL_CONFIG_TYPE) RequestValidateBusinessConfig() error {
	if self.G文件上传最大值 != nil && *self.G文件上传最大值 <= 0 {
		return errors.New("文件上传最大值必须大于 0")
	}
	return nil
}
