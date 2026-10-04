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
	Id   string `json:"id" gorm:"primaryKey;column:id;size:64"`
	Code string `json:"code" gorm:"column:code;size:64;uniqueIndex"`

	Filename string `json:"filename" gorm:"column:filename;size:256;uniqueIndex"`
	MIME     string `json:"mime" gorm:"column:mime;size:32"`
	Size     int64  `json:"size" gorm:"column:size"`
	MD5      string `json:"md5" gorm:"column:md5;size:32"`
	SHA256   string `json:"sha256" gorm:"column:sha256;size:64"`
	Width    int    `json:"width" gorm:"column:width"`
	Height   int    `json:"height" gorm:"column:height"`

	DateOnly  string    `json:"dateOnly" gorm:"column:date_only;size:16;index"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (i *Image) GetFilePath() string {
	return fmt.Sprintf("/%s/%s", i.DateOnly, i.Filename)
}
