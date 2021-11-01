package record

import (
	"Jumia_todoList/usecase/record"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase record.RecordingUseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {
	h.UseCase = useCase.(record.RecordingUseCase)

	recordRouter := router.Group("/record")
	{
		recordRouter.POST("/", h.add)
		recordRouter.PATCH("/", h.edit)
		recordRouter.DELETE("/", h.remove)
		recordRouter.POST("/filter", h.filter)

	}
}

func (h *GinHandler) add(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)
}

func (h *GinHandler) edit(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)

}

func (h *GinHandler) remove(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)

}

func (h *GinHandler) filter(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)
}
