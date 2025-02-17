package models

import (
	"time"

	"gorm.io/gorm"
)

type RoomBasic struct {
	gorm.Model
	Identity  string    `gorm:"column:identity;type:varchar(36); uniqueIndex;not null" json:"identity"`
	Name      string    `gorm:"column:name;type:varchar(100);not null" json:"name"`
	BeginAt   time.Time `gorm:"column:begin_at;type:datetime" json:"begin_at"`
	EndAt     time.Time `gorm:"column:end_at;type:datetime" json:"end_at"`
	CreatedId uint      `gorm:"column:created_id;type:int(11);not null" json:"created_id"` // creator
}

func (table *RoomBasic) TableName() string {
	return "room_basic"
}
