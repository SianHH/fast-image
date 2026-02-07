package middleware

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var respPool = sync.Pool{
	New: func() any {
		return &responseBodyWriter{
			ResponseWriter: nil,
			body:           &bytes.Buffer{},
		}
	},
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

type Data struct {
	UserAgent string
	Ip        string
	Method    string
	Host      string
	Url       string
	ReqHeader string
	ReqBody   string
	RespBody  string
	Status    int
	Ms        int64
}

func Logger(log *zap.Logger, console bool, conditions ...func(ctx *gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag = false
		for _, cond := range conditions {
			if cond == nil {
				continue
			}
			if cond(c) {
				tag = true
				break
			}
		}
		if !tag {
			return
		}
		start := time.Now()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		host := c.Request.Host
		url := c.Request.RequestURI
		ip := c.ClientIP()
		reqBody, err := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		writer := respPool.Get().(*responseBodyWriter)
		defer respPool.Put(writer)
		writer.ResponseWriter = c.Writer
		c.Writer = writer

		c.Next()

		respBody := writer.body.Bytes()
		if err == nil && len(respBody) >= 4096 {
			respBody = respBody[:4096]
		}
		writer.body.Reset() // 需要重置
		status := c.Writer.Status()
		ms := time.Now().Sub(start).Milliseconds()
		log.Info("请求日志",
			zap.String("userAgent", userAgent),
			zap.String("ip", ip),
			zap.String("method", method),
			zap.String("host", host),
			zap.String("url", url),
			zap.ByteString("reqBody", reqBody),
			zap.ByteString("respBody", respBody),
			zap.Int("status", status),
			zap.Int64("ms", ms),
		)
		if console {
			fmt.Printf(`[请求日志] =======================
[请求日志] userAgent: %s
[请求日志] ip: %s
[请求日志] method: %s
[请求日志] host: %s
[请求日志] url: %s
[请求日志] reqBody: %s
[请求日志] respBody: %s
[请求日志] status: %d
[请求日志] ms: %dms
`, userAgent, ip, method, host, url, reqBody, respBody, status, ms)
		}
	}
}
