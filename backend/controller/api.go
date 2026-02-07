package controller

import (
	"fast-image/pkg/bean"
	"fast-image/service"

	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	var req service.AuthLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Response.Param(c, err)
		return
	}
	token, err := service.Service.AuthLogin(c.ClientIP(), req)
	if err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}
	bean.Response.OkData(c, token)
}

func AuthCaptcha(c *gin.Context) {
	captcha, err := service.Service.AuthCaptcha(c.ClientIP())
	if err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}
	bean.Response.OkData(c, captcha)
}

func ImageList(c *gin.Context) {
	var req service.ImageListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Response.Param(c, err)
		return
	}
	list := service.Service.ImageList(req)
	bean.Response.OkData(c, list)
}

func ImageUpload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}

	if fileHeader.Size == 0 {
		bean.Response.Fail(c, "无法读取图片大小")
		return
	}

	if fileHeader.Size > 100*1024*1024 {
		bean.Response.Fail(c, "图片大小限制100M")
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}
	defer file.Close()
	result, err := service.Service.ImageUpload(file)
	if err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}
	bean.Response.OkData(c, result)
}

func ImageDelete(c *gin.Context) {
	var req service.ImageDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Response.Param(c, err)
		return
	}
	if err := service.Service.ImageDelete(req); err != nil {
		bean.Response.Fail(c, err.Error())
		return
	}
	bean.Response.Ok(c)
}

func ImageQuery(c *gin.Context) {
	filename := c.Param("filename")
	service.Service.ImageQuery(filename, c.Writer)
}
