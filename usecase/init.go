package usecase

import (
	"Jumia_todoList/config/logging"
	"Jumia_todoList/usecase/list"
	"Jumia_todoList/usecase/record"
	"Jumia_todoList/usecase/tag"
	"fmt"
	"github.com/rs/zerolog"
)

var ServiceBuilder DefaultService

type Service interface {
	Load(interface{}, *zerolog.Logger)
}

func init() {
	fmt.Println("START LOADING SERVICES ....")

	ServiceBuilder = NewDefaultService(logging.Log).
		Use(list.ListingName, &list.ListingService{}).
		Use(record.RecordingName, &record.RecordingService{}).
		Use(tag.TaggingName, &tag.TaggingService{})

	fmt.Println("FINISH LOADING SERVICE 🚀🚀🚀🚀")
}
