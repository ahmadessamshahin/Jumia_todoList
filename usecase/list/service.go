package list

//The actual impl of the service

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

func (l *Service) Create(i model.ListCreateInput) (int, error) {
	var o entity.List
	pkg.Cast(i, &o)
	return l.Repo.Create(&o)
}

func (l *Service) Update(i model.ListUpdateInput) error {
	var o entity.List
	pkg.Cast(i, &o)
	return l.Repo.Update(&o, i.ID)
}

func (l *Service) Delete(i model.ListRemoveInput) error {
	return l.Repo.Delete(i.ID)
}

func (l *Service) GetAllLists(i model.GetAllListInput) []entity.List {

	lists, err := l.Repo.GetAllLists(i.Limit, i.Offset)

	if err != nil {
		l.Logger.Error().
			Msgf("failed to get list with error: %s", err)
	}

	return lists

}
