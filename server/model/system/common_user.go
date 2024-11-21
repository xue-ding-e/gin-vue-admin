package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

// 这个CommonUser字段前期的业务没分开之前公共方法以及公共字段放在这里后期明确分开之后可以单独放到各自的结构体中
// TODO:Avatar和HeadrImg逻辑重复
type CommonUser struct {
	global.GVA_MODEL
	Username      string         `json:"userName" gorm:"index;comment:用户登录名"`                                                             // 用户登录名
	Password      string         `json:"-"  gorm:"comment:用户登录密码"`                                                                       // 用户登录密码
	NickName      string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                    // 用户昵称
	HeaderImg     string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`             // 用户头像
	AuthorityId   uint           `json:"authorityId" gorm:"default:5555;comment:用户角色ID"`                                                   // 用户角色ID
	Authority     SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`                      // 用户角色
	Authorities   []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`                                                     // 多用户角色
	Phone         string         `json:"phone"  gorm:"comment:用户手机号"`                                                                     // 用户手机号
	Email         string         `json:"email"  gorm:"comment:用户邮箱"`                                                                       // 用户邮箱
	Enable        int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                           //用户是否被冻结 1正常 2冻结
	OriginSetting common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置

	CreatedBy uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者"`
}

func (CommonUser) TableName() string {
	return "sys_users"
}

func (self *CommonUser) GetUsername() string {
	return self.Username
}

func (self *CommonUser) GetNickname() string {
	return self.NickName
}

func (self *CommonUser) GetUserId() uint {
	return self.ID
}

func (self *CommonUser) GetAuthorityId() uint {
	return self.AuthorityId
}

func (self *CommonUser) GetUserInfo() any {
	return *self
}
