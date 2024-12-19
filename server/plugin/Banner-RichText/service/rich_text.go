package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Banner-RichText/model/request"
)

var RichText = new(richText)

type richText struct{}

// CreateRichText 创建富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) CreateRichText(richText *model.RichText) (err error) {
	err = global.GVA_DB.Create(richText).Error
	return err
}

// DeleteRichText 删除富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) DeleteRichText(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.RichText{}, "id = ?", ID).Error
	return err
}

// DeleteRichTextByIds 批量删除富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) DeleteRichTextByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.RichText{}, "id in ?", IDs).Error
	return err
}

// UpdateRichText 更新富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) UpdateRichText(richText model.RichText) (err error) {
	err = global.GVA_DB.Model(&model.RichText{}).Where("id = ?", richText.ID).Updates(&richText).Error
	return err
}

// GetRichText 根据ID获取富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) GetRichText(ID string) (richText model.RichText, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&richText).Error
	return
}

// GetRichTextInfoList 分页获取富文本记录
// Author [yourname](https://github.com/yourname)
func (s *richText) GetRichTextInfoList(info request.RichTextSearch) (list []model.RichText, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.RichText{})
	var richTexts []model.RichText
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
	err = db.Find(&richTexts).Error
	return richTexts, total, err
}

// FindRichTextByName 根据名称获取富文本记录
func (s *richText) FindRichTextByName(name string) (richText []model.RichText, err error) {
	err = global.GVA_DB.Where("name = ?", name).Find(&richText).Error
	return
}
