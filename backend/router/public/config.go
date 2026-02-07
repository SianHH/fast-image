package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigResp struct {
	MaxPoolCount int `json:"maxPoolCount"`
}

func InitConfig(group *gin.RouterGroup) {
	group.Group("config/query", func(c *gin.Context) {
		c.JSON(http.StatusOK, ConfigResp{})
	})
}
