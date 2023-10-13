package controller

import (
	"bulebell/logic"
	"bulebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidparam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c,CodeInvalidparam,errData)
		return
	}
	userID,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeNeedLogin)
		return
	}
	if err := logic.PostVote(userID,p);err != nil{
		zap.L().Error("logic.PostVote failed",zap.Error(err))
		ResponseError(c,CodeServerBusy)
	}
	ResponseSuccess(c,nil)
}
