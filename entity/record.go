package entity

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model
	ID          int64     `gorm:"primary_key;auto_increment;not_null;" json:"id"`
	Title       string    `gorm:"type:varchar(40)" json:"title"`
	Description string    `gorm:"type:varchar(40)" json:"description"`
	Due         time.Time `gorm:"type:date" json:"due"`
	Completed   bool      `gorm:"type:boolean" json:"Completed"`
	ListRefer   uint
	Tags        []Tag `gorm:"foreignKey:RecordRefer;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
