package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

//Interfaces to what is exposing to the outer package

type ListingUseCase interface {
	Create(model.ListCreateInput) error
	Update(model.ListUpdateInput) error
	Delete(model.ListRemoveInput) error
	Get(int) (*entity.List, error)
	GetAllLists(model.GetAllListInput) ([]entity.List, error)
}
