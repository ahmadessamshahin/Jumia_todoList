package tag

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
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

func (l *TaggingService) Create(model.TagCreateInput) *entity.Tag { return nil }
func (l *TaggingService) Update(model.TagUpdateInput) *entity.Tag { return nil }
func (l *TaggingService) Delete(model.TagDeleteInput) error       { return nil }
