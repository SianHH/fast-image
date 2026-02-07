package model

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	IMAGE_PREFIX = "img-id:"
)

func GenImageID() string {
	return IMAGE_PREFIX + ulid.Make().String()
}

type Image struct {
	Id   string `json:"id"`
	Code string `json:"code"`

	Filename string `json:"filename"`
	MIME     string `json:"mime"`
	Size     int64  `json:"size"`
	MD5      string `json:"md5"`
	SHA256   string `json:"sha256"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`

	DateOnly  string    `json:"dateOnly"`
	CreatedAt time.Time `json:"createdAt"`
}

func (i *Image) GetFilePath() string {
	return fmt.Sprintf("/%s/%s", i.DateOnly, i.Filename)
}
