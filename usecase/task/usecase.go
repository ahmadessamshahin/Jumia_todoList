package task

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type UseCase interface {
	Create(model.TaskCreateInput) (int, error)
	Update(model.TaskUpdateInput) error
	Delete(model.TaskRemoveInput) error
	Filter(input model.TaskFilterInput) ([]entity.Task, error)
	Get(id int) []entity.Task
}
