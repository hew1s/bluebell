package routes

import (
	"bulebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine{
	r := gin.Default()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	r.GET("/",func (c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	return r
}