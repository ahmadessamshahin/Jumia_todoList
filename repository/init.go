package repo

import (
	"Jumia_todoList/config/constant"
	"Jumia_todoList/config/database"
	env "Jumia_todoList/config/environment"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/entity"
	listRepo "Jumia_todoList/repository/list"
	tagRepo "Jumia_todoList/repository/tag"
	taskRepo "Jumia_todoList/repository/task"

	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var RepositoryBuilder DefaultRepository

type Repository interface {
	Load(*zerolog.Logger, *gorm.DB)
}

func init() {
	fmt.Println("START LOADING REPOSITORY ....")

	RepositoryBuilder = NewDefaultRepository(logging.Log, database.Connect(env.DatabaseEnv)).
		Use(constant.ListInjectionName, &listRepo.Instance{}).
		Use(constant.TaskInjectionName, &taskRepo.Instance{}).
		Use(constant.TagInjectionName, &tagRepo.Instance{})

	fmt.Println("MIGRATING ....")

	database.Migrate(RepositoryBuilder.ORM, &entity.List{}, &entity.Task{}, &entity.Tag{})

	fmt.Println("FINISH MIGRATING ....")

	fmt.Println("FINISH Loading Repository 🚀🚀🚀")

}
