package tag

import (
	"Jumia_todoList/api/model"
)

type TaggingUseCase interface {
	Create(model.TagCreateInput) error
	Delete(output model.TagRemoveInput) error
}
