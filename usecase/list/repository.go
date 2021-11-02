package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type ListingRepository interface {
	Create(*entity.List) error
	Update(*entity.List, int) error
	Delete(int) error
	Get(int) (*entity.List, error)
	GetAllLists(model.GetAllListInput) []entity.List
}
