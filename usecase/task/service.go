package task

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

func (l *Service) Create(i model.TaskCreateInput) (int, error) {
	var o entity.Task
	pkg.Cast(i, &o)
	return l.Repo.Create(&o, i.Tags)
}

func (l *Service) Update(i model.TaskUpdateInput) error {
	var o entity.Task
	pkg.Cast(i, &o)
	return l.Repo.Update(&o, i.TaskID)
}

func (l *Service) Delete(i model.TaskRemoveInput) error {
	return l.Repo.Delete(i.ID)
}

func (l *Service) Filter(i model.TaskFilterInput) ([]entity.Task, error) {
	tasks, err := l.Repo.FilterList(i.ListId, i.Tags)

	if err != nil {
		l.Logger.Error().
			Msgf("failed to get list with error: %s", err)
	}

	return tasks, nil
}

func (l *Service) FilterAllList(i model.TaskFilterAllInput) ([]entity.Task, error) {
	tasks, err := l.Repo.FilterInAllList(i.Tags)

	if err != nil {
		l.Logger.Error().
			Msgf("failed to get list with error: %s", err)
	}

	return tasks, nil
}

func (l *Service) Get(id int) []entity.Task {
	tasks, err := l.Repo.Get(id)

	if err != nil {
		l.Logger.Error().
			Msgf("failed to get list with error: %s", err)
	}

	return tasks
}
