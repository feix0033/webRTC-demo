package models

import (
	"gorm.io/gorm"
)

type RoomUser struct {
	gorm.Model
	RoomId uint `gorm:"column:room_id;type:int(11);not null" json:"room_id"`
	UserId uint `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
}

func (table *RoomUser) TableName() string {
	return "room_user"
}
