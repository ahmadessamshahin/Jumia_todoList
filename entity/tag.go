package entity

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id     int64  `gorm:"primary_key;auto_increment;not_null;" json:"id"`
	Name   string `gorm:"type:varchar(40)" json:"name"`
	TaskID uint   `json:"task_id"`
}
