package entity


import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id      string   `gorm:"primary_key" json:"id"`
	Name    string   `gorm:"type:varchar(40)" json:"name"`
	RecordRefer   uint
}
