package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
type LoginBussinessResponse struct {
	User      system.BussinesUser `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}

// TODO 业务和系统分开
//type LoginResponse struct {
//	Token     string `json:"token"`
//	ExpiresAt int64  `json:"expiresAt"`
//}
//type LoginBussinessResponse struct {
//	User system.BussinesUser `json:"user"`
//	LoginResponse
//}
//
//type LoginSysUserResponse struct {
//	User system.SysUser `json:"user"`
//	LoginResponse
//}
