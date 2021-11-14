package tag

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"Jumia_todoList/pkg"

	"github.com/rs/zerolog"
)

type Service struct {
	Repo   Repository
	Logger *zerolog.Logger
}

func (l *Service) Load(repository interface{}, logger *zerolog.Logger) {
	l.Repo = repository.(Repository)
	l.Logger = logger
}

func (l *Service) Create(i model.TagCreateInput) (int, error) {
	var o entity.Tag
	pkg.Cast(i, &o)
	return l.Repo.Create(&o)
}
func (l *Service) Delete(i model.TagRemoveInput) error {
	return l.Repo.Delete(i.TagId)

}
