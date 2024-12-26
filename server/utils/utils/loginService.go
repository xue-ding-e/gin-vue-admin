package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemervice "github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

func GvaBaseLoginService(l *request.Login) (interface{}, error) {
	u := &system.SysUser{CommonUser: system.CommonUser{
		Username: l.Username, Password: l.Password,
	}}

	return systemervice.UserServiceApp.Login(u)
}
