package global

import (
	"fast-image/configs"
	"fast-image/pkg/jwt"
	"fast-image/pkg/orm"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var (
	Config configs.Config
	Jwt    *jwt.Tool
	Logger *zap.Logger
	Cron   *cron.Cron
	DB     orm.Interface
)
