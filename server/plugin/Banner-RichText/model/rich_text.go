
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RichText 富文本 结构体
type RichText struct {
    global.GVA_MODEL
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:;"`  //名称
    Context  *string `json:"context" form:"context" gorm:"column:context;comment:;type:text;"`  //内容
}


// TableName 富文本 RichText自定义表名 rich_text
func (RichText) TableName() string {
    return "rich_text"
}

