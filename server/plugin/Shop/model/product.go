package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Product 产品 结构体
type Product struct {
	global.GVA_MODEL
	ProdId             *int           `json:"prodId" form:"prodId" gorm:"column:prod_id;comment:产品ID;size:20;"`                                                             //产品ID
	ProdName           *string        `json:"prodName" form:"prodName" gorm:"column:prod_name;comment:商品名称;size:300;"`                                                      //商品名称
	ShopId             *int           `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`                                                             //店铺id
	OriPrice           *float64       `json:"oriPrice" form:"oriPrice" gorm:"column:ori_price;comment:原价;size:15;"`                                                         //原价
	Price              *float64       `json:"price" form:"price" gorm:"column:price;comment:现价;size:15;"`                                                                   //现价
	Brief              *string        `json:"brief" form:"brief" gorm:"column:brief;comment:简要描述,卖点等;size:500;"`                                                            //简要描述,卖点等
	Content            *string        `json:"content" form:"content" gorm:"column:content;comment:详细描述;"`                                                                   //详细描述
	Pic                *string        `json:"pic" form:"pic" gorm:"column:pic;comment:商品主图;size:255;"`                                                                      //商品主图
	Imgs               *string        `json:"imgs" form:"imgs" gorm:"column:imgs;comment:商品图片，以,分割;size:1000;"`                                                             //商品图片，以,分割
	Status             *int           `json:"status" form:"status" gorm:"column:status;comment:默认是1，表示正常状态, -1表示删除, 0下架;size:10;"`                                          //默认是1，表示正常状态, -1表示删除, 0下架
	CategoryId         *int           `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:商品分类;size:20;"`                                                 //商品分类
	SoldNum            *int           `json:"soldNum" form:"soldNum" gorm:"column:sold_num;comment:销量;size:10;"`                                                            //销量
	TotalStocks        *int           `json:"totalStocks" form:"totalStocks" gorm:"column:total_stocks;comment:总库存;size:10;"`                                               //总库存
	DeliveryMode       datatypes.JSON `json:"deliveryMode" form:"deliveryMode" gorm:"column:delivery_mode;comment:配送方式json见TransportModeVO;type:text;"swaggertype:"object"` //配送方式json见TransportModeVO
	DeliveryTemplateId *int           `json:"deliveryTemplateId" form:"deliveryTemplateId" gorm:"column:delivery_template_id;comment:运费模板id;size:19;"`                      //运费模板id

	SkuList []Sku `gorm:"-" json:"skuList"`
}

// TableName 产品 Product自定义表名 product
func (Product) TableName() string {
	return "shop_product"
}
