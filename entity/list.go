package entity

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID    int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name  string `gorm:"type:varchar(40);uniqueIndex;" json:"name"`
	Tasks []Task `gorm:"foreignKey:ListID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
