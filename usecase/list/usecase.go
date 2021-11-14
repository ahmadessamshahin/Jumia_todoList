package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

//Interfaces to what is exposing to the outer package

type UseCase interface {
	Create(model.ListCreateInput) (int, error)
	Update(model.ListUpdateInput) error
	Delete(model.ListRemoveInput) error
	GetAllLists(model.GetAllListInput) []entity.List
}
