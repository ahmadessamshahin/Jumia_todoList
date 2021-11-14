package tag

import "Jumia_todoList/entity"

type Repository interface {
	Create(*entity.Tag) (int, error)
	Delete(int) error
}
