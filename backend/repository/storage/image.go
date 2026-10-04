package storage

import (
	"errors"
	"fast-image/model"

	"gorm.io/gorm"
)

// SetImage 新增或更新一条图片记录
func SetImage(img model.Image) error {
	if img.Id == "" {
		return errors.New("image id is empty")
	}
	return db.Save(&img).Error
}

// DelImage 删除图片记录（关联的 code/filename 索引随行一并删除）
func DelImage(img model.Image) error {
	return db.Where("id = ?", img.Id).Delete(&model.Image{}).Error
}

func GetImageById(id string) (img model.Image, err error) {
	err = db.Where("id = ?", id).First(&img).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return img, errors.New("image not exist")
	}
	return img, err
}

func GetImageByCode(code string) (img model.Image, err error) {
	err = db.Where("code = ?", code).First(&img).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return img, errors.New("code not exist")
	}
	return img, err
}

func GetImageByFilename(filename string) (img model.Image, err error) {
	err = db.Where("filename = ?", filename).First(&img).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return img, errors.New("filename not exist")
	}
	return img, err
}

// ListImages 按日期范围分页查询图片记录，按创建时间倒序（新图在前）。
// startOn/endOn 为空表示对应条件不限制；offset/limit 为分页参数。
func ListImages(startOn, endOn string, offset, limit int) (list []model.Image, total int64, err error) {
	query := db.Model(&model.Image{})
	if startOn != "" {
		query = query.Where("date_only >= ?", startOn)
	}
	if endOn != "" {
		query = query.Where("date_only <= ?", endOn)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = query.Order("created_at DESC, id DESC").
		Offset(offset).Limit(limit).
		Find(&list).Error
	return list, total, err
}
