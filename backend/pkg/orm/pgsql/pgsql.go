package pgsql

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Pgsql struct {
	db  *gorm.DB
	log *lumberjack.Logger
}

func (impl *Pgsql) GetDB() *gorm.DB {
	return impl.db
}

func (impl *Pgsql) AutoMigrate(table ...any) error {
	return impl.GetDB().AutoMigrate(table...)
}

func (impl *Pgsql) Close() {
	if impl.log != nil {
		_ = impl.log.Close()
	}
	d, err := impl.db.DB()
	if err != nil {
		return
	}
	_ = d.Close()
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Prefix   string
	DbName   string
	Extend   string
}

func NewDB(config Config, logLevel string, toFile string, console bool) (*Pgsql, error) {
	var impl Pgsql
	var dns = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
		config.Host, config.Port, config.Username, config.Password, config.DbName, config.Extend)
	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭自动建表的外键约束
	}
	gormConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   config.Prefix,
		SingularTable: true,
	}

	loggerLevel := logger.Silent
	switch logLevel {
	case "info", "INFO", "Info":
		loggerLevel = logger.Info
	case "warn", "WARN", "Warn":
		loggerLevel = logger.Warn
	case "error", "ERROR", "Error":
		loggerLevel = logger.Error
	}
	logConfig := logger.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		LogLevel:                  loggerLevel, // 日志级别
		IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,       // 禁用彩色打印
	}

	var writers []io.Writer
	if console {
		writers = append(writers, os.Stdout)
	}
	if toFile != "" {
		dbLog := &lumberjack.Logger{
			Filename:   toFile,
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 10,
			LocalTime:  true,
			Compress:   true,
		}
		writers = append(writers, dbLog)
		impl.log = dbLog
	}
	gormConfig.Logger = logger.New(
		log.New(io.MultiWriter(writers...), "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logConfig,
	)
	d, err := gorm.Open(postgres.Open(dns), &gormConfig)
	if err != nil {
		return nil, err
	}
	db, err := d.DB()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 2)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	impl.db = d
	return &impl, nil
}
