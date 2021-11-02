package record

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
)

type RecordingUseCase interface {
	Create(model.RecordCreateInput) error
	Update(model.RecordUpdateInput) error
	Delete(model.RecordRemoveInput) error
	Filter(input model.RecordFilterInput) ([]entity.Record, error)
}
