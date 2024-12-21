
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Brand 品牌表 结构体
type Brand struct {
    global.GVA_MODEL
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:品牌名称;"`  //品牌名称
    Logo  *string `json:"logo" form:"logo" gorm:"column:logo;comment:品牌图标;"`  //品牌图标
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 品牌表 Brand自定义表名 shop_brand
func (Brand) TableName() string {
    return "shop_brand"
}

