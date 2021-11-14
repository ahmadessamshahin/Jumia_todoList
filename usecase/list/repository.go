package list

import (
	"Jumia_todoList/entity"
)

type Repository interface {
	Create(*entity.List) (int, error)
	Update(*entity.List, int) error
	Delete(int) error
	Get(int) (*entity.List, error)
	GetAllLists(int, int) ([]entity.List, error)
}
