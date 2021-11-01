package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type ListingRepository interface {
	Create() *entity.List
	Update() *entity.List
	Delete(string) error
	Get(string) *entity.List
	GetAllLists(model.GetAllListInput) []entity.List
}
