package bootstrap

import (
	"errors"
	"fast-image/global"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func InitServer() {
	server := &http.Server{
		Addr:    global.Config.Address,
		Handler: engine,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Release()
			global.Logger.Fatal("server listen fail", zap.Error(err))
		}
	}()
	time.Sleep(time.Second)
	releaseFunc = append(releaseFunc, func() {
		_ = server.Close()
	})
	global.Logger.Info("Web服务监听地址: " + server.Addr)
}
