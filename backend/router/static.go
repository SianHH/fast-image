package router

import (
	"archive/zip"
	"bytes"
	"fast-image/web"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var fileContentTypeMap = map[string]string{
	".js":       "application/javascript",
	".mjs":      "application/javascript",
	".css":      "text/css",
	".manifest": "text/cache-manifest",
	".png":      "image/png",
	".jpg":      "image/jpeg",
	".jpeg":     "image/jpeg",
	".svg":      "image/svg+xml",
	".ico":      "image/x-icon",
	".json":     "application/json",
	".html":     "text/html; charset=utf-8",
	".htm":      "text/html; charset=utf-8",
	".txt":      "text/plain; charset=utf-8",
	".wasm":     "application/wasm",
}

// MatchFile 更安全地匹配文件Content-Type
func MatchFile(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	if contentType, ok := fileContentTypeMap[ext]; ok {
		return contentType
	}
	return "application/octet-stream" // 更安全的默认值
}

func StaticFile(zipFile []byte, callback func(fileMap map[string][]byte)) error {
	result := make(map[string][]byte)
	zipReader, err := zip.NewReader(bytes.NewReader(zipFile), int64(len(zipFile)))
	if err != nil {
		return fmt.Errorf("加载静态资源失败: %w", err)
	}

	for _, file := range zipReader.File {
		open, err := file.Open()
		if err != nil {
			return fmt.Errorf("打开文件 %s 失败: %w", file.Name, err)
		}

		data, err := io.ReadAll(open)
		_ = open.Close()
		if err != nil {
			return fmt.Errorf("读取文件 %s 失败: %w", file.Name, err)
		}

		result[file.Name] = data
	}

	callback(result)
	return nil
}

func InitStatic(engine *gin.Engine) error {
	return StaticFile(web.Static(), func(fileMap map[string][]byte) {
		// 预检查index.html是否存在
		indexHTML, _ := fileMap["dist/index.html"]
		if len(indexHTML) == 0 {
			indexHTML = []byte("no embedded resources")
		}

		for fileKey, fileBytes := range fileMap {
			// 创建局部变量避免闭包问题
			fKey := fileKey
			fBytes := fileBytes

			ginStaticFilePath := strings.TrimPrefix(fKey, "dist/")
			if ginStaticFilePath == "" {
				continue
			}

			engine.GET(ginStaticFilePath,
				cacheControlMiddleware(),
				serveFileHandler(fKey, fBytes),
			)
		}

		engine.NoRoute(func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
		})
	})
}

// 中间件和处理器工厂函数
func cacheControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "public, max-age=86400")
		c.Next()
	}
}

func serveFileHandler(fileKey string, fileBytes []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, MatchFile(fileKey), fileBytes)
	}
}
