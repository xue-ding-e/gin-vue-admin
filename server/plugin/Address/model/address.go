package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Address 个人收货地址 结构体
type Address struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:收件人名称;"`                          //名称
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:收件人电话;"`                       //电话
	Province    *int   `json:"province" form:"province" gorm:"column:province;comment:收件省份;"`               //省份
	ProvinceStr string `json:"provinceStr" form:"provinceStr" gorm:"column:province_str;comment:收件省份（回显）;"` //省份（回显）
	City        *int   `json:"city" form:"city" gorm:"column:city;comment:收件城市;"`                           //城市
	CityStr     string `json:"cityStr" form:"cityStr" gorm:"column:city_str;comment:收件城市（回显）;"`             //城市（回显）
	Area        *int   `json:"area" form:"area" gorm:"column:area;comment:收件区;"`                            //区域
	AreaStr     string `json:"areaStr" form:"areaStr" gorm:"column:area_str;comment:收件区（回显）;"`              //区域（回显）
	Street      string `json:"street" form:"street" gorm:"column:street;comment:收件详细地址;"`                   //详细地址
	UserID      uint   `json:"userID" form:"userID" gorm:"column:user_id;comment:用户id;"`                    //用户id
	Active      *bool  `json:"active" form:"active" gorm:"column:active;comment:是否默认;"`                     //是否默认
}

// TableName 个人收货地址 Address自定义表名 address
func (Address) TableName() string {
	return "address"
}
