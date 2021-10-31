package entity

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID     string   `gorm:"primary_key" json:"id"`
	Name    string   `gorm:"type:varchar(40)" json:"name"`
	Records []Record `gorm:"foreignKey:ListRefer;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
