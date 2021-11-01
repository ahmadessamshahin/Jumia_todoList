package tag

import (
	"Jumia_todoList/usecase/tag"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase tag.TaggingUseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {
	h.UseCase = useCase.(tag.TaggingUseCase)

	tagRouter := router.Group("/tag")
	{
		tagRouter.GET("/", h.get)
		tagRouter.POST("/", h.add)
		tagRouter.DELETE("/", h.remove)
	}
}

func (h *GinHandler) add(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)

}

func (h *GinHandler) remove(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)

}

func (h *GinHandler) get(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)
}
