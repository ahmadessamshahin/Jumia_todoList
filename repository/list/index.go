package list

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

func (l *Instance) Create(i *entity.List) (int, error) {
	err := l.ORM.Save(i).Error
	return int(i.ID), err
}

func (l *Instance) Update(i *entity.List, id int) error {
	return l.ORM.Model(&entity.List{}).Where("id=?", id).Updates(i).Error
}

func (l *Instance) Delete(id int) error {
	return l.ORM.Unscoped().Delete(&entity.List{}, id).Error
}

func (l *Instance) Get(id int) (*entity.List, error) {
	list := &entity.List{}
	l.ORM.Preload("tags").Preload("tasks").Find(list, id)
	fmt.Printf("toto %#v\n", list)
	return list, nil
}

func (l *Instance) GetAllLists(limit, offset int) (lists []entity.List, err error) {
	err = l.ORM.Limit(limit).Offset(offset).
		Find(&lists).Error
	return
}

func (l *Instance) ClearTable() {
	l.ORM.Exec("DELETE FROM lists")
	l.ORM.Exec("ALTER SEQUENCE lists_id_seq RESTART WITH 1")
}
