package service

import (
	"fast-image/global"
	"fast-image/model"
	"fast-image/repository/storage"
	"fmt"
	"slices"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type ImageListReq struct {
	StartOn string `json:"startOn"`
	EndOn   string `json:"endOn"`
}

type ImageItem struct {
	Id   string `json:"id"`
	Code string `json:"code"`

	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	MD5      string `json:"md5"`
	SHA256   string `json:"sha256"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`

	DateOnly  string `json:"dateOnly"`
	CreatedAt string `json:"createdAt"`
}

func (s *service) ImageList(req ImageListReq) (list []ImageItem) {
	prefix := []byte(model.IMAGE_PREFIX)
	_ = global.BadgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			fmt.Printf("KEY=%q\n", it.Item().Key())
			var image model.Image
			if err := storage.GetStruct(txn, string(it.Item().Key()), &image); err != nil {
				continue
			}

			if image.DateOnly < req.StartOn || image.DateOnly > req.EndOn {
				continue
			}

			i := ImageItem{
				Id:        image.Id,
				Code:      image.Code,
				Filename:  image.Filename,
				Size:      image.Size,
				MD5:       image.MD5,
				SHA256:    image.SHA256,
				Width:     image.Width,
				Height:    image.Height,
				DateOnly:  image.DateOnly,
				CreatedAt: image.CreatedAt.Format(time.DateTime),
			}

			list = append(list, i)
		}
		return nil
	})
	slices.Reverse(list)
	return list
}
