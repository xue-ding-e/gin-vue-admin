package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
)

var Good = new(good)

type good struct{}

// CreateGood 创建商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) CreateGood(good *shop.Good) (err error) {
	err = global.GVA_DB.Create(good).Error
	return err
}

// DeleteGood 删除商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) DeleteGood(ID string) (err error) {
	err = global.GVA_DB.Delete(&shop.Good{}, "id = ?", ID).Error
	return err
}

// DeleteGoodByIds 批量删除商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) DeleteGoodByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]shop.Good{}, "id in ?", IDs).Error
	return err
}

// UpdateGood 更新商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) UpdateGood(good shop.Good) (err error) {
	err = global.GVA_DB.Model(&shop.Good{}).Where("id = ?", good.ID).Updates(&good).Error
	return err
}

// GetGood 根据ID获取商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) GetGood(ID string) (good shop.Good, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&good).Error
	return
}

// GetGoodInfoList 分页获取商品记录
// Author [yourname](https://github.com/yourname)
func (s *good) GetGoodInfoList(info request.GoodSearch) (list []shop.Good, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shop.Good{})
	var goods []shop.Good
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&goods).Error
	return goods, total, err
}

func (s *good) GetGoodPublic() {

}
