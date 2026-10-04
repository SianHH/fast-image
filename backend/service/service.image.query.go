package service

import (
	"fast-image/global"
	"fast-image/repository/storage"
	"io"
	"os"
)

func (s *service) ImageQuery(filename string, writer io.Writer) {
	img, err := storage.GetImageByFilename(filename)
	if err != nil {
		return
	}

	file, err := os.OpenFile(global.GetBasePath()+"/data/images/"+img.GetFilePath(), os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	_, _ = io.Copy(writer, file)
}
