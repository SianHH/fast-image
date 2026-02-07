package bootstrap

import (
	"fast-image/global"
	"fast-image/pkg/logger"
)

func InitLogger() {
	loggerFile, loggerLevel := global.GetLogger()
	mode := global.GetAppMode()
	global.Logger = logger.NewLogger(func() logger.Option {
		var to = []string{loggerFile}
		if mode == global.APP_MODE_DEV {
			to = append(to, "console")
		}
		return logger.Option{
			To: to,
			WriterOption: logger.WriterOption{
				Level:      loggerLevel,
				MaxSize:    100,
				MaxAge:     30,
				MaxBackups: 10,
				Compress:   true,
			},
		}
	}())

	releaseFunc = append(releaseFunc, func() {
		_ = global.Logger.Sync()
	})
}
