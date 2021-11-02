package tag

import (
	"Jumia_todoList/api/model"
	"github.com/rs/zerolog"
)

const TaggingName string = "TAG"

type TaggingService struct {
	Repo   TaggingRepository
	Logger *zerolog.Logger
}

func (l *TaggingService) Load(repository interface{}, logger *zerolog.Logger) {
	l.Repo = repository.(TaggingRepository)
	l.Logger = logger
}

func (l *TaggingService) Create(model.TagCreateInput) error { return nil }
func (l *TaggingService) Delete(model.TagRemoveInput) error { return nil }
