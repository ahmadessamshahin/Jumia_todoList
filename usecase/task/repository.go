package task

import (
	"Jumia_todoList/entity"
)

type Repository interface {
	Create(*entity.Task, []string) (int, error)
	Update(*entity.Task, uint) error
	Delete(int) error
	FilterList(int, []string) ([]entity.Task, error)
	Get(int) ([]entity.Task, error)
	FilterInAllList([]string) ([]entity.Task, error)
}
