package orm

import (
	"gorm.io/gorm"
)

type Interface interface {
	GetDB() *gorm.DB
	AutoMigrate(tables ...any) error
	Close()
}
