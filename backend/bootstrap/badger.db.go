package bootstrap

import (
	"context"
	"fast-image/global"
	"fast-image/repository/storage"
	"time"

	"github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"
)

func InitBadgerDB() {
	options := badger.DefaultOptions(global.GetBasePath() + "/data/badger").
		WithValueLogFileSize(64 << 20). // 64MB
		WithIndexCacheSize(2 << 20).    // 2MB
		WithBlockCacheSize(4 << 20).    // 默认 100MB -> 4MB
		WithMemTableSize(16 << 20).     // 默认 64MB -> 16MB
		WithNumMemtables(2).            // 默认 5 -> 2
		WithNumLevelZeroTables(2).      // 默认 5 -> 2
		WithNumLevelZeroTablesStall(4). // 默认 15 -> 4
		WithValueThreshold(1024)        // value全部使用vlog，降低内存表压力

	var err error
	global.BadgerDB, err = badger.Open(options)
	if err != nil {
		Release()
		global.Logger.Fatal("init badgerDB fail", zap.Any("config", global.Config), zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(time.Minute * 10)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_ = global.BadgerDB.RunValueLogGC(0.5)
			case <-ctx.Done():
				return
			}
		}
	}()
	storage.SetBadgerDB(global.BadgerDB)
	releaseFunc = append(releaseFunc, func() {
		cancel()
		_ = global.BadgerDB.Close()
	})
}
