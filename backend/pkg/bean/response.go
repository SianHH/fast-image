package bean

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	OK            = 0 // 成功
	FAIL          = 1 // 失败
	AUTH_INVALID  = 2 // 登录失效
	AUTH_NO_LOGIN = 3 // 未登录
	AUTH_NO_ALLOW = 4 // 不允许
)

type response struct{}

var Response = response{}

func (response) Ok(c *gin.Context) {
	c.JSON(200, gin.H{"code": OK})
}

func (response) OkData(c *gin.Context, data interface{}) {
	result := gin.H{
		"code": OK,
	}
	if data != nil {
		result["data"] = data
	}
	c.JSON(200, result)
}

func (response) Param(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"code": FAIL,
		"msg":  validateError(err),
	})
}

func (response) Fail(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": FAIL,
		"msg":  msg,
	})
}

func (response) AuthInvalid(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": AUTH_INVALID,
		"msg":  "登录失效",
	})
}

func (response) AuthDeny(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"code": AUTH_INVALID,
		"msg":  msg,
	})
}

func (response) AuthNoLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": AUTH_NO_LOGIN,
		"msg":  "未登录",
	})
}

func (response) AuthNoAllow(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": AUTH_NO_ALLOW,
		"msg":  "权限不足",
	})
}

func validateError(err error) string {
	var er validator.ValidationErrors
	if errors.As(err, &er) {
		field := er[0].Field()
		if field == er[0].StructField() {
			field = strings.ToLower(field[0:1]) + field[1:]
		}
		switch er[0].Tag() {
		case "required":
			return field + "不能为空"
		case "min":
			if er[0].Type().String() == "string" {
				return field + "不能小于" + er[0].Param() + "位"
			}
			return field + "不能小于" + er[0].Param()
		}
		return field + "错误"
	} else {
		return err.Error()
	}
}
