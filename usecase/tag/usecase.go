package tag

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type TaggingUseCase interface {
	Create(model.TagCreateInput) *entity.Tag
	Update(model.TagUpdateInput) *entity.Tag
	Delete(model.TagDeleteInput) error
}
