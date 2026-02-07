package storage

import (
	"errors"
	"fast-image/model"

	"github.com/dgraph-io/badger/v4"
)

func SetImage(txn *badger.Txn, img model.Image) error {
	if err := SetStruct(txn, img.Id, img, 0); err != nil {
		return err
	}
	if err := SetString(txn, img.Code, img.Id, 0); err != nil {
		return err
	}
	if err := SetString(txn, img.Filename, img.Id, 0); err != nil {
		return err
	}
	return nil
}

func DelImage(txn *badger.Txn, img model.Image) error {
	if err := txn.Delete([]byte(img.Id)); err != nil {
		return err
	}
	if err := txn.Delete([]byte(img.Code)); err != nil {
		return err
	}
	return nil
}

func GetImageById(txn *badger.Txn, id string) (img model.Image, err error) {
	if err := GetStruct(txn, id, &img); err != nil {
		return img, err
	}
	return img, nil
}

func GetImageByCode(txn *badger.Txn, code string) (img model.Image, err error) {
	id := GetString(txn, code)
	if id == "" {
		return img, errors.New("code not exist")
	}
	if err := GetStruct(txn, id, &img); err != nil {
		return img, err
	}
	return img, nil
}

func GetImageByFilename(txn *badger.Txn, filename string) (img model.Image, err error) {
	id := GetString(txn, filename)
	if id == "" {
		return img, errors.New("code not exist")
	}
	if err := GetStruct(txn, id, &img); err != nil {
		return img, err
	}
	return img, nil
}
