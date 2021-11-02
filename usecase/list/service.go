package list

//The actual impl of the service

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"github.com/rs/zerolog"
)

const ListingName string = "LIST"

type ListingService struct {
	Repo   ListingRepository
	Logger *zerolog.Logger
}

func (l *ListingService) Load(repository interface{}, logger *zerolog.Logger) {
	l.Repo = repository.(ListingRepository)
	l.Logger = logger
}

func (l *ListingService) Create(i model.ListCreateInput) error {
	var o entity.List
	helper.Cast(i, &o)
	return l.Repo.Create(&o)
}
func (l *ListingService) Update(i model.ListUpdateInput) error {
	var o entity.List
	helper.Cast(i, &o)
	return l.Repo.Update(&o, i.ID)
}
func (l *ListingService) Delete(i model.ListRemoveInput) error {
	return l.Repo.Delete(i.ID)
}
func (l *ListingService) Get(id int) (*entity.List, error) {
	l.Repo.Get(id)
	return nil, nil
}
func (l *ListingService) GetAllLists(model.GetAllListInput) ([]entity.List, error) { return nil, nil }
