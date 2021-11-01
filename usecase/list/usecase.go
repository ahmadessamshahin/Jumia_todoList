package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

//Interfaces to what is exposing to the outer package

type ListingUseCase interface {
	Create() *entity.List
	Update(model.ListUpdateInput) *entity.List
	Delete(string) error
	Get(string) *entity.List
	GetAllLists(model.GetAllListInput) []entity.List
}
