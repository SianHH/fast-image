package storage

import (
	"bytes"
	"fast-image/pkg/utils"
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v4"
)

const (
	ip_security_key = "security:ip:"
	captcha_key     = "captcha:"
)

func SetIpSecurity(ip string, security bool) {
	_ = db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(fmt.Sprintf("%s%s", ip_security_key, ip)), utils.TrinaryOperation(security, []byte{1}, []byte{2}))
		_ = txn.SetEntry(entry)
		return nil
	})
}

func GetIpSecurity(ip string) (result bool) {
	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("%s%s", ip_security_key, ip)))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		result = !bytes.Equal(val, []byte{2})
		return nil
	})
	return result
}

func SetCaptcha(key string, value string, duration time.Duration) {
	_ = db.Update(func(txn *badger.Txn) error {
		k := []byte(fmt.Sprintf("%s%s", captcha_key, key))
		entry := badger.NewEntry(k, []byte(value)).WithTTL(duration)
		_ = txn.SetEntry(entry)
		return nil
	})
}

func ValidCaptcha(key string, target string) (result bool) {
	_ = db.Update(func(txn *badger.Txn) error {
		k := []byte(fmt.Sprintf("%s%s", captcha_key, key))
		item, err := txn.Get(k)
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		result = bytes.Equal(val, []byte(target))
		_ = txn.Delete(k)
		return nil
	})
	return result
}
