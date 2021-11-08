package usecase

import (
	"Jumia_todoList/config/constant"
	"Jumia_todoList/config/logging"
	repo "Jumia_todoList/repository"
	"Jumia_todoList/usecase/list"
	"Jumia_todoList/usecase/tag"
	"Jumia_todoList/usecase/task"
	"fmt"
	"github.com/rs/zerolog"
)

var ServiceBuilder DefaultService

type Service interface {
	Load(interface{}, *zerolog.Logger)
}

func init() {
	fmt.Println("START LOADING SERVICES ....")

	ServiceBuilder = NewDefaultService(logging.Log, repo.RepositoryBuilder).
		Use(constant.ListInjectionName, &list.Service{}).
		Use(constant.TaskInjectionName, &task.Service{}).
		Use(constant.TagInjectionName, &tag.Service{})

	fmt.Println("FINISH LOADING SERVICE 🚀🚀🚀🚀")
}
