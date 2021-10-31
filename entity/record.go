package entity

import (
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model
	ID          string    `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(40)" json:"title"`
	Description string    `gorm:"type:varchar(40)" json:"description"`
	Due         time.Time `gorm:"type:date" json:"due"`
	Completed   bool      `gorm:"type:boolean" json:"Completed"`
	ListRefer   uint
	Tags []Tag `gorm:"foreignKey:RecordRefer;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
