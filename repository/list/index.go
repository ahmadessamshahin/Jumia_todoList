package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type ListingInstance struct {
	Logger *zerolog.Logger
	ORM    *gorm.DB
}

func (l *ListingInstance) Load(logger *zerolog.Logger, orm *gorm.DB) {
	l.ORM = orm
	l.Logger = logger
}

func (l *ListingInstance) Create(i *entity.List) error {
	return l.ORM.Create(i).Error
}

func (l *ListingInstance) Update(i *entity.List, id int) error {
	return l.ORM.Model(&entity.List{}).Where("id=?", id).Updates(i).Error
}

func (l *ListingInstance) Delete(id int) error {
	return l.ORM.Delete(&entity.List{}, id).Error

}

func (l *ListingInstance) Get(id int) (*entity.List, error) {
	list := &entity.List{}
	err := l.ORM.Find(list, id).Error

	if err != nil {
		return nil, err
	}

	records := entity.Record{}
	err = l.ORM.Model(list).Where("id = ?", id).Association("Records").Find(&records)

	if err != nil {
		return nil, err
	}

	return list, err
}

func (l *ListingInstance) GetAllLists(model.GetAllListInput) []entity.List {
	return nil
}
