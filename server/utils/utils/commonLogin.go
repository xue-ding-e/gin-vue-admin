package utils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

func commonLogin(c *gin.Context,
	loginFunc func(l *request.Login) (user interface{}, err error),
	tokenNext func(c *gin.Context, user interface{}) (data interface{}, err error),
) (interface{}, error) {
	var l request.Login

	if err := c.ShouldBindJSON(&l); err != nil {
		return nil, err
	}
	key := c.ClientIP()
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		return nil, err
	}
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	oc := openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || (l.CaptchaId != "" && l.Captcha != "" && base64Captcha.DefaultMemStore.Verify(l.CaptchaId, l.Captcha, true)) {
		user, err := loginFunc(&l)
		if err != nil {
			global.GVA_LOG.Debug("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			return nil, errors.New("用户名不存在或者密码错误")
		}
		//可以更换这里的签发token的函数
		if data, err := tokenNext(c, user); err != nil {
			global.GVA_LOG.Debug("签发token失败!", zap.Error(err))
			return nil, errors.New("签发token失败")
		} else {
			return data, nil
		}
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	return nil, errors.New("验证码错误")
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
