package service

import (
	"errors"
	"fast-image/global"
	"fast-image/repository/storage"
	"time"
)

type AuthLoginReq struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	CaptchaKey   string `json:"captchaKey"`
	CaptchaValue string `json:"captchaValue"`
}

func (s *service) AuthLogin(ip string, req AuthLoginReq) (token string, err error) {
	defer func() {
		if err != nil {
			storage.SetIpSecurity(ip, false)
		} else {
			storage.SetIpSecurity(ip, true)
		}
	}()

	if !storage.GetIpSecurity(ip) && !storage.ValidCaptcha(req.CaptchaKey, req.CaptchaValue) {
		return "", errors.New("验证码错误")
	}

	if req.Account != global.Config.Account || req.Password != global.Config.Password {
		return "", errors.New("账号密码错误")
	}

	token, err = global.Jwt.GenerateToken(global.Jwt.NewClaims("", map[string]string{}, time.Hour*24*60))
	return token, err
}
