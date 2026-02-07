package service

import (
	"fast-image/global"
	"fast-image/repository/storage"
	"io"
	"os"

	"github.com/dgraph-io/badger/v4"
)

func (s *service) ImageQuery(filename string, writer io.Writer) {
	_ = global.BadgerDB.View(func(txn *badger.Txn) error {
		img, err := storage.GetImageByFilename(txn, filename)
		if err != nil {
			return nil
		}

		file, err := os.OpenFile(global.GetBasePath()+"/data/images/"+img.GetFilePath(), os.O_RDONLY, 0644)
		if err != nil {
			return nil
		}
		defer file.Close()

		_, _ = io.Copy(writer, file)
		return nil
	})
}
