package bootstrap

import (
	"fast-image/global"

	"github.com/robfig/cron/v3"
)

var TaskFunc func()

func InitTask() {
	global.Cron = cron.New()
	if TaskFunc != nil {
		TaskFunc()
	}
	global.Cron.Start()
	global.Logger.Info("初始化TASK完成")
}
