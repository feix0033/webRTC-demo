package service

import (
	"time"
	"webRTC-demo/internal/helper"
	"webRTC-demo/internal/models"

	"github.com/gin-gonic/gin"
)

func MeetingCreate(c *gin.Context) {
	in := new(MeetingCreateRequest)

	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "parameters error",
		})
		return
	}

	err = models.DB.Create(&models.RoomBasic{
		Identity:  helper.GetUUID(),
		Name:      "",
		BeginAt:   time.UnixMilli(int64(in.CreateAt)),
		EndAt:     time.UnixMilli(int64(in.EndAt)),
		CreatedId: 0, // todo get user id from auth middleware
	}).Error
	if err != nil {
		c.JSON(503, gin.H{
			"code":    503,
			"message": "internal error",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "meeting create",
	})
}
