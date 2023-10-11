package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code":10000,
	"msg":msg,
	"data":{}
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(), //通过状态码返回不同的msg信息
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode,msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(), //通过状态码返回不同的msg信息
		Data: data,
	})
}

