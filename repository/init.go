package repo

import (
	"Jumia_todoList/config/database"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/entity"
	"Jumia_todoList/usecase/list"
	"Jumia_todoList/usecase/record"
	"Jumia_todoList/usecase/tag"

	listingRepo "Jumia_todoList/repository/list"
	recordingRepo "Jumia_todoList/repository/record"
	taggingRepo "Jumia_todoList/repository/tag"

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

	RepositoryBuilder = NewDefaultRepository(logging.Log, database.Connect()).
		Use(list.ListingName, &listingRepo.ListingInstance{}).
		Use(record.RecordingName, &recordingRepo.RecordingInstance{}).
		Use(tag.TaggingName, &taggingRepo.TaggingInstance{})

	fmt.Println("MIGRATING ....")

	database.Migrate(RepositoryBuilder.ORM, &entity.List{}, &entity.Record{}, &entity.Tag{})

	fmt.Println("FINISH MIGRATING ....")

	fmt.Println("FINISH Loading Repository 🚀🚀🚀")

}
