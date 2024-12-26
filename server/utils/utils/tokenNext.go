package utils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func GvaBaseTokenNext(c *gin.Context, u interface{}) (interface{}, error) {
	user, ok := u.(*system.SysUser)
	if !ok {
		return systemRes.LoginSysUserResponse{}, errors.New("user 断言为 system.SysUser 失败")
	}
	token, claims, err := utils.LoginToken(user)
	if err != nil {
		// 这些地方尽量不要用error级别以防被刷Log日志
		global.GVA_LOG.Debug("获取token失败!", zap.Error(err))
		return systemRes.LoginSysUserResponse{}, errors.New("获取token失败")
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		return systemRes.LoginSysUserResponse{
			LoginResponse: systemRes.LoginResponse{
				Token:     token,
				ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			},
			User: *user,
		}, nil
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Debug("设置登录状态失败!", zap.Error(err))
			return systemRes.LoginSysUserResponse{}, errors.New("设置登录状态失败")
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		return systemRes.LoginSysUserResponse{
			User: *user,
			LoginResponse: systemRes.LoginResponse{
				Token:     token,
				ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			},
		}, nil
	} else if err != nil {
		global.GVA_LOG.Debug("设置登录状态失败!", zap.Error(err))
		return systemRes.LoginSysUserResponse{}, errors.New("设置登录状态失败")
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			return systemRes.LoginSysUserResponse{}, errors.New("jwt作废失败")
		}
		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
			return systemRes.LoginSysUserResponse{}, errors.New("设置登录状态失败")
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		return systemRes.LoginSysUserResponse{
			User: *user,
			LoginResponse: systemRes.LoginResponse{
				Token:     token,
				ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
			},
		}, nil
	}
}
