package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)
const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// 获取当前用户登录的ID
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}


func getPageInfo(c *gin.Context) (int64,int64) {
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")
	var (
		pageNum int64
		pageSize int64
		err error
	)
	pageNum,err = strconv.ParseInt(pageNumStr,10,64)
	if err != nil {
		pageNum = 1
	}
	pageSize,err = strconv.ParseInt(pageSizeStr,10,64)
	if err != nil {
		pageSize = 10
	}
	return pageNum,pageSize
}