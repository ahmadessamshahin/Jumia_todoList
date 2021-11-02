package repo

import (
	"fmt"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type DefaultRepository struct {
	Logger      *zerolog.Logger
	RepoMethods map[string]Repository
	ORM         *gorm.DB
}

func NewDefaultRepository(Logger *zerolog.Logger, orm *gorm.DB) DefaultRepository {
	return DefaultRepository{
		Logger:      Logger,
		ORM:         orm,
		RepoMethods: make(map[string]Repository),
	}
}

func (s DefaultRepository) Use(method string, factory Repository) DefaultRepository {
	factory.Load(s.Logger, s.ORM)
	s.RepoMethods[method] = factory
	return s
}

func (s DefaultRepository) Resolve(repoType string) interface{} {
	if method, ok := s.RepoMethods[repoType]; ok {
		return method
	}
	panic(fmt.Errorf("THIS REPOSITORY TYPE: %s DOESN'T EXIST", repoType))
}
