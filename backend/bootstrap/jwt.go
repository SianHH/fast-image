package bootstrap

import (
	"fast-image/global"
	"fast-image/pkg/jwt"
	"fast-image/pkg/utils"
)

func InitJwt() {
	salt := utils.RandStr(32, utils.AllDict)
	global.Jwt = jwt.NewTool(salt)
}
