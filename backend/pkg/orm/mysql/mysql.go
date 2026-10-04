package mysql

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

type Mysql struct {
	db  *gorm.DB
	log *lumberjack.Logger
}

func (impl *Mysql) GetDB() *gorm.DB {
	return impl.db
}

func (impl *Mysql) AutoMigrate(table ...any) error {
	return impl.GetDB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table...)
}

func (impl *Mysql) Close() {
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
	Extend   string
	DbName   string
}

func NewDB(config Config, logLevel string, toFile string, console bool) (*Mysql, error) {
	var impl Mysql
	var dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", config.Username, config.Password, config.Host, config.Port, config.DbName, config.Extend)
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
	d, err := gorm.Open(mysql2.Open(dns), &gormConfig)
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
