package bootstrap

import (
	"context"
	"fast-image/global"
	"fast-image/model"
	"fast-image/pkg/orm/sqlite"
	"fast-image/repository/storage"
	"time"

	"go.uber.org/zap"
)

// InitDB 初始化 SQLite 数据库（替代原 BadgerDB）
func InitDB() {
	mode := global.GetAppMode()
	_, logLevel := global.GetLogger()

	db, err := sqlite.NewDB(
		global.GetBasePath()+"/data/app.db",
		logLevel,
		"",
		mode == global.APP_MODE_DEV,
	)
	if err != nil {
		Release()
		global.Logger.Fatal("init sqlite db fail", zap.Any("config", global.Config), zap.Error(err))
	}
	global.DB = db
	storage.SetDB(global.DB.GetDB())

	// 自动建表
	if err := global.DB.AutoMigrate(
		&model.Image{},
		&storage.IpSecurity{},
		&storage.Captcha{},
	); err != nil {
		Release()
		global.Logger.Fatal("auto migrate sqlite fail", zap.Any("config", global.Config), zap.Error(err))
	}

	// 定时清理过期验证码
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				storage.CleanExpiredCaptcha()
			case <-ctx.Done():
				return
			}
		}
	}()
	releaseFunc = append(releaseFunc, func() {
		cancel()
		global.DB.Close()
	})
}
