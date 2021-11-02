package record

import (
	"Jumia_todoList/entity"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type RecordingInstance struct {
	Logger *zerolog.Logger
	ORM    *gorm.DB
}

func (l *RecordingInstance) Load(logger *zerolog.Logger, orm *gorm.DB) {
	l.ORM = orm
	l.Logger = logger
}

func (l *RecordingInstance) Create() *entity.Record {
	return nil
}

func (l *RecordingInstance) Update() *entity.Record {
	return nil
}

func (l *RecordingInstance) Delete() error {
	return nil
}

func (l *RecordingInstance) FilterList() *entity.Record {
	return nil
}
func (l *RecordingInstance) FilterInAllList() *entity.Record {
	return nil
}
