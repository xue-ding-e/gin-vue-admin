
package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
    "gorm.io/gorm"
)

var Brand = new(brand)

type brand struct {}
// CreateBrand 创建品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) CreateBrand(brand *model.Brand) (err error) {
	err = global.GVA_DB.Create(brand).Error
	return err
}

// DeleteBrand 删除品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) DeleteBrand(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Brand{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&model.Brand{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteBrandByIds 批量删除品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) DeleteBrandByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&model.Brand{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&model.Brand{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateBrand 更新品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) UpdateBrand(brand model.Brand) (err error) {
	err = global.GVA_DB.Model(&model.Brand{}).Where("id = ?",brand.ID).Updates(&brand).Error
	return err
}

// GetBrand 根据ID获取品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) GetBrand(ID string) (brand model.Brand, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&brand).Error
	return
}

// GetBrandInfoList 分页获取品牌表记录
// Author [yourname](https://github.com/yourname)
func (self *brand) GetBrandInfoList(info request.BrandSearch) (list []model.Brand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Brand{})
    var brands []model.Brand
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&brands).Error
	return  brands, total, err
}

func (self *brand)GetBrandPublic() {

}