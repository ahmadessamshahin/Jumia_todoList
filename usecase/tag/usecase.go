package tag

import (
	"Jumia_todoList/api/model"
)

type UseCase interface {
	Create(model.TagCreateInput) error
	Delete(model.TagRemoveInput) error
}
