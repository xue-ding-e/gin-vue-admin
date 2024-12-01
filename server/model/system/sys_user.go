package system

type SysUser struct {
	CommonUser
}

func (SysUser) TableName() string {
	return "sys_users"
}
