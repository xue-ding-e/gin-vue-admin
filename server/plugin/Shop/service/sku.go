
package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
)

var Sku = new(sku)

type sku struct {}
// CreateSku 创建sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) CreateSku(sku *model.Sku) (err error) {
	err = global.GVA_DB.Create(sku).Error
	return err
}

// DeleteSku 删除sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) DeleteSku(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Sku{},"id = ?",ID).Error
	return err
}

// DeleteSkuByIds 批量删除sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) DeleteSkuByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Sku{},"id in ?",IDs).Error
	return err
}

// UpdateSku 更新sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) UpdateSku(sku model.Sku) (err error) {
	err = global.GVA_DB.Model(&model.Sku{}).Where("id = ?",sku.ID).Updates(&sku).Error
	return err
}

// GetSku 根据ID获取sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) GetSku(ID string) (sku model.Sku, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sku).Error
	return
}

// GetSkuInfoList 分页获取sku表记录
// Author [yourname](https://github.com/yourname)
func (self *sku) GetSkuInfoList(info request.SkuSearch) (list []model.Sku, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Sku{})
    var skus []model.Sku
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
	err = db.Find(&skus).Error
	return  skus, total, err
}

func (self *sku)GetSkuPublic() {

}