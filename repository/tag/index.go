package tag

import (
	"Jumia_todoList/entity"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type TaggingInstance struct {
	Logger *zerolog.Logger
	ORM    *gorm.DB
}

func (l *TaggingInstance) Load(logger *zerolog.Logger, orm *gorm.DB) {
	l.ORM = orm
	l.Logger = logger
}

func (l *TaggingInstance) Create() *entity.Tag {
	return nil
}

func (l *TaggingInstance) Delete() error {
	return nil
}
