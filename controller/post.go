package controller

import (
	"bulebell/logic"
	"bulebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	// 1.参数
	if err := c.ShouldBindJSON(p);err !=nil{
		zap.L().Error("CreatePost param failed",zap.Error(err))
		ResponseError(c,CodeInvalidparam)
		return
	}
	// 2.创建
	userID ,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeNeedLogin)
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p);err != nil{
		zap.L().Error("logic.CreatePost failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
	// 3.返回响应
}