package service

import (
	"encoding/base64"
	"errors"
	"fast-image/pkg/captcha"
	"fast-image/pkg/utils"
	"fast-image/repository/storage"
	"time"

	"github.com/fatedier/golib/log"
	"go.uber.org/zap"
)

type AuthCaptchaResp struct {
	Key      string `json:"key"`
	Bs64     string `json:"bs64"`
	Security bool   `json:"security"`
}

func (s *service) AuthCaptcha(ip string) (result AuthCaptchaResp, err error) {
	code := utils.RandStr(4, utils.NumDict)
	bytes, err := captcha.Generate(120, 40, code)
	if err != nil {
		log.Error("生成图片验证码失败", zap.Error(err))
		return result, errors.New("获取失败")
	}
	result.Bs64 = base64.StdEncoding.EncodeToString(bytes)
	result.Key = utils.RandStr(16, utils.AllDict)
	storage.SetCaptcha(result.Key, code, time.Minute*5)
	result.Security = storage.GetIpSecurity(ip)
	return result, nil
}
