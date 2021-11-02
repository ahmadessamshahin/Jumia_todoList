package record

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"github.com/rs/zerolog"
)

const RecordingName string = "RECORD"

type RecordingService struct {
	Repo   RecordingRepository
	Logger *zerolog.Logger
}

func (l *RecordingService) Load(repository interface{}, logger *zerolog.Logger) {
	l.Repo = repository.(RecordingRepository)
	l.Logger = logger
}

func (l *RecordingService) Create(model.RecordCreateInput) error { return nil }
func (l *RecordingService) Update(model.RecordUpdateInput) error { return nil }
func (l *RecordingService) Delete(model.RecordRemoveInput) error { return nil }
func (l *RecordingService) Filter(model.RecordFilterInput) ([]entity.Record, error) {
	return nil, nil
}
