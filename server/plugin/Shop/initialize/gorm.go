package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	shop "github.com/flipped-aurora/gin-vue-admin/server/plugin/Shop/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(shop.Good{}, shop.Category{}, shop.Sku{})
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
