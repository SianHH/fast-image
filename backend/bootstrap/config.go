package bootstrap

import (
	"fast-image/configs"
	"fast-image/global"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func InitConfig() {
	path := global.GetBasePath()
	configFilePath := path + "/data/config.yaml"
	_ = os.MkdirAll(filepath.Dir(configFilePath), 0666)

	global.Logger.Info("初始化配置文件", zap.String("path", configFilePath))
	stat, err := os.Stat(configFilePath)
	if err != nil {
		global.Logger.Warn("配置文件不存在", zap.String("path", configFilePath), zap.Error(err))
		global.Config = defaultConfig
		global.Config.LoadFromEnv()
		if err := global.Config.StoreToFile(configFilePath); err != nil {
			global.Logger.Fatal("持久化默认配置文件失败", zap.String("path", configFilePath), zap.Error(err))
		} else {
			global.Logger.Info("初始化配置文件完成", zap.String("path", configFilePath))
		}
		return
	}
	if stat.IsDir() {
		global.Logger.Fatal("配置文件是目录，不是文件", zap.String("path", configFilePath), zap.String("path", configFilePath))
	}
	if err := global.Config.LoadFromFile(configFilePath); err != nil {
		global.Logger.Fatal("读取配置文件失败", zap.String("path", configFilePath), zap.Error(err))
	}
	global.Config.LoadFromEnv()
	global.Logger.Info("配置文件加载完成", zap.String("path", configFilePath))
}

var defaultConfig = configs.Config{
	Address:  "0.0.0.0:8080",
	Account:  "admin",
	Password: "admin",
}
