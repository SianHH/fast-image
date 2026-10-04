// migrate 一次性工具：将旧版 BadgerDB 中的图片元数据迁移到 SQLite。
//
// 用法（迁移前请先停止后端服务，且 badger 目录未被其它进程占用）：
//
//	cd backend/cmd/migrate
//	go mod tidy
//	go run . --badger ../../data/badger --sqlite ../../data/app.db
//
// 说明：
//   - 只迁移 "img-id:" 前缀的图片记录；
//   - 建表直接复用 gormlite.AutoMigrate（与应用内 model.Image 完全一致的 DDL），
//     避免应用启动时 AutoMigrate 因 DDL 不一致而重建表导致数据丢失；
//   - 写入使用主键覆盖（Save），重复执行不会报错；
//   - 迁移完成后，可删除 data/badger 目录（旧 KV 数据不再使用）。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	badger "github.com/dgraph-io/badger/v4"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// imageRow 的字段与 GORM 标签必须与主模块 model.Image 保持一致
type imageRow struct {
	Id        string    `json:"id" gorm:"primaryKey;column:id;size:64"`
	Code      string    `json:"code" gorm:"column:code;size:64;uniqueIndex"`
	Filename  string    `json:"filename" gorm:"column:filename;size:256;uniqueIndex"`
	MIME      string    `json:"mime" gorm:"column:mime;size:32"`
	Size      int64     `json:"size" gorm:"column:size"`
	MD5       string    `json:"md5" gorm:"column:md5;size:32"`
	SHA256    string    `json:"sha256" gorm:"column:sha256;size:64"`
	Width     int       `json:"width" gorm:"column:width"`
	Height    int       `json:"height" gorm:"column:height"`
	DateOnly  string    `json:"dateOnly" gorm:"column:date_only;size:16;index"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

// TableName 必须与主模块 model.Image 一致（SingularTable 命名 -> image）
func (imageRow) TableName() string { return "image" }

const imagePrefix = "img-id:"

func main() {
	badgerDir := flag.String("badger", "", "BadgerDB 数据目录，如 data/badger")
	sqliteFile := flag.String("sqlite", "", "SQLite 数据库文件，如 data/app.db")
	flag.Parse()

	if *badgerDir == "" || *sqliteFile == "" {
		fmt.Println("用法: migrate --badger <badger-dir> --sqlite <sqlite-file>")
		os.Exit(1)
	}

	// 打开 BadgerDB（Windows 下 badger 不支持只读模式，需确保目标目录未被占用）
	opts := badger.DefaultOptions(*badgerDir)
	bd, err := badger.Open(opts)
	if err != nil {
		fmt.Printf("打开 BadgerDB 失败: %v\n", err)
		os.Exit(1)
	}
	defer bd.Close()

	// 打开 SQLite 并建表（DDL 与应用 AutoMigrate 完全一致）
	_ = os.MkdirAll(filepath.Dir(*sqliteFile), 0755)
	db, err := gorm.Open(gormlite.Open(*sqliteFile), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("打开 SQLite 失败: %v\n", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&imageRow{}); err != nil {
		fmt.Printf("初始化 SQLite 表失败: %v\n", err)
		os.Exit(1)
	}

	// 扫描 BadgerDB 中全部图片记录并写入 SQLite
	count := 0
	err = bd.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		prefixBytes := []byte(imagePrefix)
		for it.Seek(prefixBytes); it.ValidForPrefix(prefixBytes); it.Next() {
			item := it.Item()
			key := string(item.Key())

			var val []byte
			if v, verr := item.ValueCopy(nil); verr == nil {
				val = v
			}
			if len(val) == 0 {
				continue
			}

			var row imageRow
			if err := json.Unmarshal(val, &row); err != nil {
				fmt.Printf("跳过损坏记录 key=%s: %v\n", key, err)
				continue
			}
			if row.Id == "" {
				row.Id = key
			}

			// Save 按主键覆盖写入，兼容重复执行
			if err := db.Save(&row).Error; err != nil {
				fmt.Printf("写入失败 key=%s: %v\n", key, err)
				continue
			}
			count++
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历 BadgerDB 失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("迁移完成，共写入 %d 条图片记录 -> %s\n", count, *sqliteFile)
}
