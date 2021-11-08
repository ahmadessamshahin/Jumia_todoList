package handler

import (
	"Jumia_todoList/api/handler/list"
	"Jumia_todoList/api/handler/tag"
	"Jumia_todoList/api/handler/task"
	"Jumia_todoList/config/constant"
	"Jumia_todoList/config/logging"
	"Jumia_todoList/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RoutingGroup interface {
	Routes(*gin.Engine, interface{})
}

var Handler *DefaultHandler

func init() {
	fmt.Println("Start Loading Router ......")
	Handler = NewDefaultHandler(logging.Log, usecase.ServiceBuilder).
		Use(constant.ListInjectionName, &list.GinHandler{}).
		Use(constant.TaskInjectionName, &task.GinHandler{}).
		Use(constant.TagInjectionName, &tag.GinHandler{})

	fmt.Println("Finish Loading Router 🚀🚀🚀🚀🚀")
}
