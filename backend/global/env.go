package global

import (
	"os"
	"path/filepath"
)

type AppMode string

const (
	APP_MODE_DEV  AppMode = "dev"
	APP_MODE_PROD AppMode = "prod"
)

var (
	base_path = ""

	logger_level = "info"
	logger_file  = ""

	app_mode AppMode = APP_MODE_PROD

	version = "1.1.0"
)

func init() {
	executable, _ := os.Executable()
	base_path = filepath.Dir(executable)
	logger_level = "info"
	logger_file = base_path + "/data/logs/fast-image.log"
}

func SetBasePath(path string) {
	if path == "" {
		return
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return
	}
	base_path = abs
	logger_file = base_path + "/data/logs/fast-image.log"
}

func GetBasePath() string {
	return base_path
}

func SetLoggerLevel(level string) {
	logger_level = level
}

func GetLogger() (loggerFile string, loggerLevel string) {
	return logger_file, logger_level
}

func SetAppMode(mode AppMode) {
	app_mode = mode
}

func GetAppMode() AppMode {
	return app_mode
}

func SetVersion(v string) {
	version = v
}

func GetVersion() string {
	return version
}
