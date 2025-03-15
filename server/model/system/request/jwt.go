package request

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

// CustomClaims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID          uint
	Username    string
	Nickname    string
	AuthorityId uint
}
