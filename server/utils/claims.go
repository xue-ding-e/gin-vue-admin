package utils

import (
	"net"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gofiber/fiber/v2"
)

// 增加cookie x-token 向来源的web添加
func ClearToken(c *fiber.Ctx) {
	// 由于 Fiber 不提供直接的 c.Request.Host，需要使用 c.Hostname() 获取域名。
	// 若包含端口，可自行解析；以下示例假设 c.Hostname() 返回的已去除端口，或无需区分端口。
	host := c.Hostname()
	if net.ParseIP(host) != nil {
		// IP请求,设置为空域名
		c.Cookie(&fiber.Cookie{Name: "x-token", Value: "", Path: "/", MaxAge: -1, HTTPOnly: false})
	} else {
		//域名请求 Cookie 域设置为 host
		c.Cookie(&fiber.Cookie{Name: "x-token", Value: "", Path: "/", Domain: host, MaxAge: -1, HTTPOnly: false})
	}
}

func SetToken(c *fiber.Ctx, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host := c.Hostname()
	if net.ParseIP(host) != nil {
		c.Cookie(&fiber.Cookie{Name: "x-token", Value: token, Path: "/", MaxAge: maxAge, HTTPOnly: false})
	} else {
		c.Cookie(&fiber.Cookie{Name: "x-token", Value: token, Path: "/", Domain: host, MaxAge: maxAge, HTTPOnly: false})
	}
}

func GetToken(c *fiber.Ctx) string {
	token := c.Cookies("x-token")
	if token == "" {
		j := NewJWT()
		token = c.Get("x-token")
		claims, err := j.ParseToken(token)
		if err != nil {
			// 不建议采用error级别,线上环境会被恶意增加日志
			global.GVA_LOG.Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构")
			return token
		}
		SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

func GetClaims(c *fiber.Ctx) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *fiber.Ctx) uint {
	if claims := c.Locals("claims"); claims == nil {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *fiber.Ctx) uint {
	if claims := c.Locals("claims"); claims == nil {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *fiber.Ctx) *systemReq.CustomClaims {
	if claims := c.Locals("claims"); claims == nil {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *fiber.Ctx) string {
	if claims := c.Locals("claims"); claims == nil {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.Username
	}
}

func LoginToken(user system.Login) (token string, claims systemReq.CustomClaims, err error) {
	j := &JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims = j.CreateClaims(systemReq.BaseClaims{
		ID:          user.GetUserId(),
		NickName:    user.GetNickname(),
		Username:    user.GetUsername(),
		AuthorityId: user.GetAuthorityId(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return
	}
	return
}
