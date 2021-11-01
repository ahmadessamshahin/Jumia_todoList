package usecase

import (
	repo "Jumia_todoList/repository"
	"fmt"
	"github.com/rs/zerolog"
)

type DefaultService struct {
	Logger         *zerolog.Logger
	serviceMethods map[string]Service
}

func NewDefaultService(Logger *zerolog.Logger) DefaultService {
	return DefaultService{
		Logger:         Logger,
		serviceMethods: make(map[string]Service),
	}
}

func (s DefaultService) Use(method string, factory Service) DefaultService {
	factory.Load(repo.RepositoryBuilder.Resolve(method), s.Logger)
	s.serviceMethods[method] = factory
	return s
}

func (s DefaultService) Resolve(serviceType string) interface{} {
	if method, ok := s.serviceMethods[serviceType]; ok {
		return method
	}
	panic(fmt.Errorf("THIS %s DOESN'T EXIST", serviceType))
}
