package service

import (
	"fast-image/global"
	"fast-image/repository/storage"
	"os"
)

type ImageDeleteReq struct {
	Id  string   `json:"id"`
	Ids []string `json:"ids"`
}

func (s *service) ImageDelete(req ImageDeleteReq) error {
	if req.Id != "" {
		req.Ids = append(req.Ids, req.Id)
	}

	for _, id := range req.Ids {
		image, err := storage.GetImageById(id)
		if err != nil {
			continue
		}
		if err := storage.DelImage(image); err != nil {
			continue
		}
		_ = os.Remove(global.GetBasePath() + "/data/images/" + image.GetFilePath())
	}
	return nil
}
