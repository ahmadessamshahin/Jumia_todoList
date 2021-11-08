package entity

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID          int64     `gorm:"primary_key;auto_increment;not_null;" json:"id" deepcopier:"field:id"`
	Title       string    `gorm:"type:varchar(40)" json:"title" deepcopier:"field:title"`
	Description string    `gorm:"type:varchar(40)" json:"description" deepcopier:"field:description"`
	Due         time.Time `gorm:"type:date" json:"due" deepcopier:"field:due"`
	Completed   bool      `gorm:"type:boolean" json:"completed" deepcopier:"field:completed"`
	ListID      uint      `json:"list_id"`
	Tags        []Tag     `gorm:"foreignKey:TaskID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags" deepcopier:"field:tags"`
}
