package storage

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB 注入 SQLite 数据库连接（由 bootstrap 在启动时调用）
func SetDB(d *gorm.DB) {
	db = d
}
