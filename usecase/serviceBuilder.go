package usecase

import (
	repo "Jumia_todoList/repository"
	"fmt"
	"github.com/rs/zerolog"
)

type DefaultService struct {
	RepositoryBuilder repo.DefaultRepository
	Logger            *zerolog.Logger
	serviceMethods    map[string]Service
}

func NewDefaultService(Logger *zerolog.Logger, repositoryBuilder repo.DefaultRepository) DefaultService {
	return DefaultService{
		RepositoryBuilder: repositoryBuilder,
		Logger:            Logger,
		serviceMethods:    make(map[string]Service),
	}
}

func (s DefaultService) Use(method string, factory Service) DefaultService {
	factory.Load(s.RepositoryBuilder.Resolve(method), s.Logger)
	s.serviceMethods[method] = factory
	return s
}

func (s DefaultService) Resolve(serviceType string) interface{} {
	if method, ok := s.serviceMethods[serviceType]; ok {
		return method
	}
	panic(fmt.Errorf("THIS %s DOESN'T EXIST", serviceType))
}
