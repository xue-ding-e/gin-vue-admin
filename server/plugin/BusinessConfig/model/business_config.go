package model

// BusinessConfig 业务设置 结构体
type BusinessConfig struct {
	Name  string `gorm:"primarykey"` // 业务名称，例如“业务1”
	Value string `gorm:"type:text"`  // 配置内容，JSON 格式，使用 text 类型以便无限制存储

}

// TableName 业务设置 BusinessConfig自定义表名 business_config
func (BusinessConfig) TableName() string {
	return "business_config"
}
