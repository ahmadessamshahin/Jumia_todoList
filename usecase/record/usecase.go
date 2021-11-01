package record

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type RecordingUseCase interface {
	Create(model.RecordCreateInput) *entity.Record
	Update(model.RecordUpdateInput) *entity.Record
	Delete(model.RecordDeleteInput) error
}
