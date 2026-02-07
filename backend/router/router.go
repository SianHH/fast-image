package router

import (
	"fast-image/bootstrap"
	"fast-image/controller"
	"fast-image/global"
	"fast-image/pkg/bean"
	"fast-image/pkg/jwt"
	"fast-image/pkg/logger"
	"fast-image/pkg/middleware"
	"fast-image/router/public"
	"net/http"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		if token == "" {
			bean.Response.AuthNoLogin(c)
			c.Abort()
			return
		}
		claims, err := global.Jwt.ValidToken(token)
		if err != nil {
			bean.Response.AuthInvalid(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func GetToken(c *gin.Context) string {
	return c.GetHeader("token")
}

func GetClaims(c *gin.Context) jwt.Claims {
	claims, exists := c.Get("claims")
	if !exists {
		panic("get claims fail")
	}
	return claims.(jwt.Claims)
}

func init() {
	bootstrap.Route = func(engine *gin.Engine) {
		// 使用GZIP压缩数据
		engine.Use(gzip.Gzip(gzip.DefaultCompression))

		// 记录HTTP请求日志
		httpLog := logger.NewLogger(logger.Option{
			To: []string{global.GetBasePath() + "/data/http_logs/log.log"},
			WriterOption: logger.WriterOption{
				Level:      "debug",
				MaxSize:    100,
				MaxAge:     30,
				MaxBackups: 10,
				Compress:   true,
			},
		})

		engine.Use(middleware.Logger(httpLog, global.GetAppMode() == global.APP_MODE_DEV, func(c *gin.Context) bool {
			if strings.HasPrefix(c.Request.RequestURI, "/api/v1/image/upload") {
				return false
			}
			if strings.HasPrefix(c.Request.RequestURI, "/api/v1/image") && c.Request.Method == http.MethodGet {
				return false
			}
			return strings.HasPrefix(c.Request.RequestURI, "/api/v")
		}))

		_ = InitStatic(engine)

		api := engine.Group("api/v1")

		authGroup := api.Group("auth")
		authGroup.POST("login", controller.AuthLogin)
		authGroup.POST("captcha", controller.AuthCaptcha)

		tunnelGroup := api.Group("image", Auth())
		tunnelGroup.POST("list", controller.ImageList)
		tunnelGroup.POST("delete", controller.ImageDelete)
		tunnelGroup.POST("upload", controller.ImageUpload)

		api.GET("image/:filename", cacheControlMiddleware(), controller.ImageQuery)

		public.InitConfig(api.Group("public"))
	}
}
