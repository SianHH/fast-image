package service

import (
	"fast-image/global"
	"fast-image/repository/storage"
	"os"

	"github.com/dgraph-io/badger/v4"
)

type ImageDeleteReq struct {
	Id  string   `json:"id"`
	Ids []string `json:"ids"`
}

func (s *service) ImageDelete(req ImageDeleteReq) error {
	if err := global.BadgerDB.Update(func(txn *badger.Txn) error {
		if req.Id != "" {
			req.Ids = append(req.Ids, req.Id)
		}

		for _, id := range req.Ids {
			image, err := storage.GetImageById(txn, id)
			if err != nil {
				continue
			}
			if err := storage.DelImage(txn, image); err != nil {
				continue
			}
			_ = os.Remove(global.GetBasePath() + "/data/images/" + image.GetFilePath())
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
