package list

//The actual impl of the service

import (
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

func (l *ListingService) Create() *entity.List                            { return nil }
func (l *ListingService) Update(model.ListUpdateInput) *entity.List       { return nil }
func (l *ListingService) Delete(string) error                             { return nil }
func (l *ListingService) Get(string) *entity.List                         { return nil }
func (l *ListingService) GetAllLists(model.GetAllListInput) []entity.List { return nil }
