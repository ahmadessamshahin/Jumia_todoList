package helper

import "gorm.io/gorm"

type FnOptions func(tx *gorm.DB) (*gorm.DB, error)
