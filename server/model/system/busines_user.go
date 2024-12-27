package system

type BusinessUser struct {
	CommonUser
}

func (BusinessUser) TableName() string {
	return "sys_users"
}
