package tag

import (
	"Jumia_todoList/entity"
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Instance struct {
	Logger *zerolog.Logger
	ORM    *gorm.DB
}

func (l *Instance) Load(logger *zerolog.Logger, orm *gorm.DB) {
	l.ORM = orm
	l.Logger = logger
}

func (l *Instance) Create(i *entity.Tag) (int, error) {
	err := l.ORM.Save(i).Error
	fmt.Println(i.ID, err)
	return int(i.ID), err
}

func (l *Instance) Delete(id int) error {
	return l.ORM.Unscoped().Delete(&entity.List{}, id).Error
}
