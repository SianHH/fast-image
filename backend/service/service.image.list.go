package service

import (
	"fast-image/pkg/bean"
	"fast-image/repository/storage"
	"time"
)

type ImageListReq struct {
	StartOn string `json:"startOn"`
	EndOn   string `json:"endOn"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
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

// ImageList 分页查询图片记录。
// startOn/endOn 非必填：为空表示对应边界不限制；page/size 为空时使用默认分页。
func (s *service) ImageList(req ImageListReq) (any, error) {
	page := bean.PageParam{Page: req.Page, Size: req.Size}
	limit := page.GetLimit()
	offset := page.GetOffset()

	images, total, err := storage.ListImages(req.StartOn, req.EndOn, offset, limit)
	if err != nil {
		return nil, err
	}

	list := make([]ImageItem, 0, len(images))
	for _, image := range images {
		list = append(list, ImageItem{
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
		})
	}
	return bean.NewPage(list, total), nil
}
