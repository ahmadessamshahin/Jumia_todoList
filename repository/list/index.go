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

func (l *ListingInstance) Create() *entity.List {
	return nil
}

func (l *ListingInstance) Update() *entity.List {
	return nil
}

func (l *ListingInstance) Delete(string) error {
	return nil
}

func (l *ListingInstance) Get(string) *entity.List {
	return nil
}

func (l *ListingInstance) GetAllLists(model.GetAllListInput) []entity.List {
	return nil
}
