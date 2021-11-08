package tag

import "Jumia_todoList/entity"

type Repository interface {
	Create(*entity.Tag) error
	Delete(int) error
}
