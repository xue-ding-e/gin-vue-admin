package utils

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

func RegisterApis(apis []system.SysApi, checkEveryOne ...bool) {
	shouldCheckEveryOne := false
	if len(checkEveryOne) > 0 {
		shouldCheckEveryOne = checkEveryOne[0]
	}

	if shouldCheckEveryOne {
		// 对每个 API 单独处理
		for i := range apis {
			var count int64
			global.GVA_DB.Model(&system.SysApi{}).Where("path = ?", apis[i].Path).Count(&count)

			if count == 0 {
				// 如果不存在，则创建新的
				if err := global.GVA_DB.Create(&apis[i]).Error; err != nil {
					return
				}
			} else {
				// 如果存在，可以选择更新或跳过，这里选择跳过
				continue
			}
		}
	} else {
		// 检查是否有任何一个 API 的 Path 已存在
		var count int64
		var apiPaths []string
		for _, api := range apis {
			apiPaths = append(apiPaths, api.Path)
		}
		global.GVA_DB.Model(&system.SysApi{}).Where("path IN (?)", apiPaths).Count(&count)
		if count > 0 {
			// 如果有任何一个存在，直接返回
			return
		}
		// 如果都不存在，批量创建
		if err := global.GVA_DB.Create(&apis).Error; err != nil {
			return
		}
	}
}

func RegisterMenus(menus []system.SysBaseMenu, isMount ...bool) {
	// 处理默认值
	shouldMount := false
	if len(isMount) > 0 {
		shouldMount = isMount[0]
	}

	parentMenu := menus[0]
	otherMenus := menus[1:]

	// 检查父菜单是否存在
	err := global.GVA_DB.Where("name = ?", parentMenu.Name).First(&parentMenu).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 父菜单不存在，创建它
		if err := global.GVA_DB.Create(&parentMenu).Error; err != nil {
			global.GVA_LOG.Sugar().Error("注册父菜单%s失败", parentMenu.Name, err.Error())
			return
		}
	} else if !shouldMount {
		// 父菜单存在且不需要强制挂载，直接返回
		return
	}

	// 检查每个子菜单
	for i := range otherMenus {
		var existCount int64
		global.GVA_DB.Model(&system.SysBaseMenu{}).Where("name = ? AND parent_id = ?", otherMenus[i].Name, parentMenu.ID).Count(&existCount)
		if existCount == 0 {
			// 子菜单不存在，设置父 ID 并创建
			otherMenus[i].ParentId = parentMenu.ID
			if err := global.GVA_DB.Create(&otherMenus[i]).Error; err != nil {
				global.GVA_LOG.Sugar().Error("注册子菜单%s失败", otherMenus[i].Name, err.Error())
				return
			}
		}
	}
}
