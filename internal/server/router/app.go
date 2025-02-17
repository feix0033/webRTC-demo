package router

import (
	"webRTC-demo/internal/server/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// metting
	r.POST("/meeting/create", service.MeetingCreate)
	return r
}
