package storage

import (
	"encoding/json"
	"time"

	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB

func SetBadgerDB(badgerDB *badger.DB) {
	db = badgerDB
}

func GetString(txn *badger.Txn, key string) string {
	item, err := txn.Get([]byte(key))
	if err != nil {
		return ""
	}
	val, err := item.ValueCopy(nil)
	if err != nil {
		return ""
	}
	return string(val)
}

func SetString(txn *badger.Txn, key string, value string, expire time.Duration) error {
	entry := badger.NewEntry([]byte(key), []byte(value))
	if expire > 0 {
		entry = entry.WithTTL(expire)
	}
	return txn.SetEntry(entry)
}

func GetStruct(txn *badger.Txn, key string, value any) error {
	item, err := txn.Get([]byte(key))
	if err != nil {
		return err
	}
	val, err := item.ValueCopy(nil)
	if err != nil {
		return err
	}
	return json.Unmarshal(val, value)
}

func SetStruct(txn *badger.Txn, key string, value any, expire time.Duration) error {
	marshal, err := json.Marshal(value)
	if err != nil {
		return err
	}
	entry := badger.NewEntry([]byte(key), marshal)
	if expire > 0 {
		entry = entry.WithTTL(expire)
	}
	return txn.SetEntry(entry)
}

func ScanByPrefix[T any](txn *badger.Txn, prefix string, filter func(data T) bool) (list []T) {
	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false
	var prefixBytes = []byte(prefix)

	it := txn.NewIterator(opts)
	defer it.Close()
	for it.Seek(prefixBytes); it.ValidForPrefix(prefixBytes); it.Next() {
		var data T
		if err := GetStruct(txn, string(it.Item().Key()), &data); err != nil {
			continue
		}
		if !filter(data) {
			continue
		}
		list = append(list, data)
	}
	return list
}
