package service

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fast-image/global"
	"fast-image/model"
	"fast-image/repository/storage"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/gen2brain/webp"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

var allowedImageMimes = map[string]bool{
	"image/jpeg": true, // 覆盖 jpg / jpeg
	"image/png":  true,
	"image/gif":  true,
	"image/webp": true,
}

func WebPQualityBySize(w, h int) int {
	pixels := w * h

	switch {
	case pixels <= 500_000: // <= 0.5MP
		return 85
	case pixels <= 2_000_000: // <= 2MP
		return 80
	case pixels <= 6_000_000: // <= 6MP
		return 75
	default:
		return 70
	}
}

func ConvertReaderToWebPWithInfo(reader io.Reader, imageInfo *ImageInfo) (io.Reader, error) {
	if imageInfo.MIME == "image/webp" || imageInfo.Format == "webp" || imageInfo.Format == "image/gif" || imageInfo.Format == "gif" {
		return reader, nil
	}

	// 必须完整解码（物理限制）
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		if err := webp.Encode(pw, img, webp.Options{
			Lossless: false,
			Quality:  WebPQualityBySize(imageInfo.Width, imageInfo.Height),
		}); err != nil {
			_ = pw.CloseWithError(err)
		}
	}()

	imageInfo.MIME = "image/webp"
	imageInfo.Format = "webp"
	return pr, nil
}

func (s *service) ImageUpload(r io.Reader) (result model.Image, err error) {
	img := model.Image{
		Id:        model.GenImageID(),
		Code:      uuid.NewString(),
		Filename:  "",
		MIME:      "",
		Size:      0,
		MD5:       "",
		SHA256:    "",
		Width:     0,
		Height:    0,
		DateOnly:  time.Now().Format(time.DateOnly),
		CreatedAt: time.Now(),
	}

	imageInfo, reader, err := GetImageInfo(r)
	if err != nil {
		return result, err
	}

	if !allowedImageMimes[imageInfo.MIME] {
		return result, errors.New("仅支持jpg/jpeg/png/webp/gif格式")
	}

	reader, err = ConvertReaderToWebPWithInfo(reader, imageInfo)
	if err != nil {
		return result, err
	}

	img.Width = imageInfo.Width
	img.Height = imageInfo.Height
	img.MIME = imageInfo.MIME
	img.Filename = fmt.Sprintf("%s_%d.%s", uuid.NewString(), time.Now().UnixMilli(), imageInfo.Format)

	md5Hash := md5.New()
	sha256Hash := sha256.New()

	filePath := global.GetBasePath() + "/data/images/" + img.GetFilePath()
	_ = os.MkdirAll(path.Dir(filePath), 0755)

	writer, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return result, err
	}
	defer writer.Close()

	mw := io.MultiWriter(writer, md5Hash, sha256Hash)
	size, err := io.Copy(mw, reader)
	if err != nil {
		return result, err
	}

	img.Size = size
	img.MD5 = hex.EncodeToString(md5Hash.Sum(nil))
	img.SHA256 = hex.EncodeToString(sha256Hash.Sum(nil))

	if err := global.BadgerDB.Update(func(txn *badger.Txn) error {
		return storage.SetImage(txn, img)
	}); err != nil {
		return result, err
	}
	return img, nil
}

type ImageInfo struct {
	Format string // jpeg / jpg / png / webp / gif
	MIME   string // image/jpeg
	Width  int
	Height int
}

func GetImageInfo(r io.Reader) (*ImageInfo, io.Reader, error) {
	const sniffSize = 512
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	head := make([]byte, sniffSize)
	n, err := io.ReadFull(tee, head)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, nil, err
	}
	head = head[:n]

	mime := http.DetectContentType(head)
	if !strings.HasPrefix(mime, "image/") {
		return nil, nil, errors.New("不是图片类型的文件")
	}

	cfg, format, err := image.DecodeConfig(io.MultiReader(bytes.NewReader(head), tee))
	if err != nil {
		return nil, nil, err
	}

	info := &ImageInfo{
		Format: format,
		MIME:   mime,
		Width:  cfg.Width,
		Height: cfg.Height,
	}

	finalReader := io.MultiReader(&buf, r)

	return info, finalReader, nil
}
