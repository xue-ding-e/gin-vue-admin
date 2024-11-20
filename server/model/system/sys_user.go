package system

type Login interface {
	GetUsername() string
	GetNickname() string
	GetUserId() uint
	GetAuthorityId() uint
	GetUserInfo() any
}

var _ Login = new(SysUser)

type SysUser struct {
	CommonUser
}

func (SysUser) TableName() string {
	return "sys_users"
}

func (s *SysUser) GetUsername() string {
	return s.Username
}

func (s *SysUser) GetNickname() string {
	return s.NickName
}

func (s *SysUser) GetUserId() uint {
	return s.ID
}

func (s *SysUser) GetAuthorityId() uint {
	return s.AuthorityId
}

func (s *SysUser) GetUserInfo() any {
	return *s
}
