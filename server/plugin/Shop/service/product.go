package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model/request"
)

var Product = new(product)

type product struct{}

// CreateProduct 创建产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) CreateProduct(product *model.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) DeleteProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Product{}, "id = ?", ID).Error
	return err
}

// DeleteProductByIds 批量删除产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) DeleteProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Product{}, "id in ?", IDs).Error
	return err
}

// UpdateProduct 更新产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) UpdateProduct(product model.Product) (err error) {
	err = global.GVA_DB.Model(&model.Product{}).Where("id = ?", product.ID).Updates(&product).Error
	return err
}

// GetProduct 根据ID获取产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) GetProduct(ID string) (product model.Product, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&product).Error
	return
}

// GetProductInfoList 分页获取产品记录
// Author [yourname](https://github.com/yourname)
func (self *product) GetProductInfoList(info request.ProductSearch) (list []model.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Product{})
	var products []model.Product
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
	err = db.Find(&products).Error
	return products, total, err
}

// 根据分类 ID 分页获取商品列表
func (self *product) GetProductsByCategoryID(categoryID uint64, page, size int) ([]model.Product, error) {
	var products []model.Product
	offset := (page - 1) * size
	err := global.GVA_DB.Table("prod p").
		Select("p.*, sd.shop_name").
		Joins("LEFT JOIN shop_detail sd ON p.shop_id = sd.shop_id").
		Where("p.category_id = ? AND p.status = 1", categoryID).
		Order("p.putaway_time DESC").
		Offset(offset).Limit(size).
		Find(&products).Error
	return products, err
}

// 根据商品 ID 获取商品详情
func (self *product) GetProductByID(prodID uint64) (*model.Product, error) {
	var product_data model.Product
	// 获取商品基本信息
	err := global.GVA_DB.Table("prod p").
		Select("p.*, sd.shop_name").
		Joins("LEFT JOIN shop_detail sd ON p.shop_id = sd.shop_id").
		Where("p.prod_id = ?", prodID).
		First(&product_data).Error
	if err != nil {
		return nil, err
	}
	return &product_data, nil
}

// 获取 SKU 列表
func (self *product) GetSkuListByProdID(prodID uint64) ([]model.Sku, error) {
	var skuList []model.Sku
	err := global.GVA_DB.Where("prod_id = ? AND is_delete = 0 AND status = 1", prodID).Find(&skuList).Error
	return skuList, err
}
