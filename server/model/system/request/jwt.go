package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/Wx/model/request"
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
