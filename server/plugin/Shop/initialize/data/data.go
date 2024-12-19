package data

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initGood struct{}

func init() {
	system.RegisterInit(int(system.GetInitCounter()), &initGood{})
}

func (i initGood) InitializerName() string {
	return shop.Good{}.TableName()
}

func (i *initGood) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&shop.Good{})
}
func (i *initGood) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&shop.Good{})
}

func (i *initGood) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	var entities []shop.Good
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, shop.Good{}.TableName()+" 表数据初始化失败!")
	}
	//next := context.WithValue(ctx, i.InitializerName(), entities)
	//return next, nil
	return ctx, nil
}

func (i *initGood) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.First(&shop.Good{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
