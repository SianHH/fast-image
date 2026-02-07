package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option struct {
	To []string // 目标文件,console会输出至stdout
	WriterOption
}

type WriterOption struct {
	Level      string // 日志等级 debug,info,warn,error,panic,fatal Default:info
	MaxSize    int    // 日志大小限制，单位MB
	MaxAge     int    // 历史日志文件保留天数
	MaxBackups int    // 最大保留历史日志数量
	Compress   bool   // 历史日志文件压缩标识
}

var DefaultOption = func() Option {
	return Option{
		To: []string{"console"},
		WriterOption: WriterOption{
			Level:      "debug",
			MaxSize:    5,
			MaxAge:     30,
			MaxBackups: 5,
			Compress:   true,
		},
	}
}

func NewLogger(option Option) *zap.Logger {
	// 日志级别
	atomicLevel := zap.NewAtomicLevel()
	switch option.Level {
	case "debug", "DEBUG":
		atomicLevel.SetLevel(zapcore.DebugLevel)
	case "info", "INFO":
		atomicLevel.SetLevel(zapcore.InfoLevel)
	case "warn", "WARN":
		atomicLevel.SetLevel(zapcore.WarnLevel)
	case "error", "ERROR":
		atomicLevel.SetLevel(zapcore.ErrorLevel)
	case "panic", "PANIC":
		atomicLevel.SetLevel(zapcore.PanicLevel)
	case "fatal", "FATAL":
		atomicLevel.SetLevel(zapcore.FatalLevel)
	default:
		atomicLevel.SetLevel(zapcore.InfoLevel)
	}
	encoderConfig := newEncoderConfig(false)
	colorEncoderConfig := newEncoderConfig(true)

	devEnCoderCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(colorEncoderConfig),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	)
	var coreList []zapcore.Core
	for _, to := range option.To {
		switch to {
		case "console":
			coreList = append(coreList, devEnCoderCore)
		default:
			writer := NewWriter(to, option)
			coreList = append(coreList, zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(writer),
				atomicLevel,
			))
		}
	}
	// 添加默认日志
	if len(coreList) == 0 {
		coreList = append(coreList, devEnCoderCore)
	}
	return zap.New(zapcore.NewTee(coreList...), zap.AddCaller())
}

func newEncoderConfig(color bool) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "line",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel: func() func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			if color {
				return zapcore.CapitalColorLevelEncoder
			}
			return zapcore.CapitalLevelEncoder
		}(),
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func NewWriter(to string, option Option) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   to,
		MaxSize:    option.MaxSize,
		MaxAge:     option.MaxAge,
		MaxBackups: option.MaxBackups,
		LocalTime:  true,
		Compress:   option.Compress,
	}
}
