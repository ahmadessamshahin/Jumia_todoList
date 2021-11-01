package handler

import (
	"Jumia_todoList/api/handler/list"
	"Jumia_todoList/api/handler/record"
	"Jumia_todoList/api/handler/tag"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/usecase"
	listingUseCase "Jumia_todoList/usecase/list"
	recordingUseCase "Jumia_todoList/usecase/record"
	taggingUseCase "Jumia_todoList/usecase/tag"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoutingGroup interface {
	Routes(*gin.Engine, interface{})
}

var Handler *DefaultHandler

func init() {
	fmt.Println("Start Loading Router ......")
	Handler = NewDefaultHandler(logging.Log).
		Use(&list.GinHandler{}, usecase.ServiceBuilder.Resolve(listingUseCase.ListingName)).
		Use(&record.GinHandler{}, usecase.ServiceBuilder.Resolve(recordingUseCase.RecordingName)).
		Use(&tag.GinHandler{}, usecase.ServiceBuilder.Resolve(taggingUseCase.TaggingName))

	fmt.Println("Finish Loading Router 🚀🚀🚀🚀🚀")
}
