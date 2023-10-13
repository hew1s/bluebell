package controller

import (
	"bulebell/logic"
	"bulebell/models"
	"strconv"

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

func GetPostHandler(c *gin.Context){
	// 1.获取参数
	pidStr := c.Param("id")
	pid,err := strconv.ParseInt(pidStr,10,64)
	if err != nil {
		zap.L().Error("get post detail with invalid param",zap.Error(err))
		ResponseError(c,CodeInvalidparam)
		return
	}
	data,err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,data)
}

// 获取帖子列表
func GetPostListHandler(c *gin.Context)  {
	//获取数据返回响应
	pageNum,pageSize := getPageInfo(c)
	data,err := logic.GetPostList(pageNum,pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList failed",zap.Error(err))
	}
	ResponseSuccess(c,data)
}