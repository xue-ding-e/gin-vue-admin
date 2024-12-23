package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (self *GVA_MODEL) FindById(db *gorm.DB, out interface{}) error {
	return db.First(out, self.ID).Error
}

func (self *GVA_MODEL) FindListById(db *gorm.DB, out interface{}) error {
	return db.Find(out, self.ID).Error
}
