package global

import (
	"fast-image/configs"
	"fast-image/pkg/jwt"

	"github.com/dgraph-io/badger/v4"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var (
	Config   configs.Config
	Jwt      *jwt.Tool
	Logger   *zap.Logger
	BadgerDB *badger.DB
	Cron     *cron.Cron
)
