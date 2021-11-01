package list

import (
	"Jumia_todoList/usecase/list"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase list.ListingUseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {

	h.UseCase = useCase.(list.ListingUseCase)

	listRouter := router.Group("/list")
	{
		listRouter.GET("/", h.get)
		listRouter.POST("/", h.add)
		listRouter.PATCH("/", h.edit)
		listRouter.DELETE("/", h.remove)
		listRouter.GET("/all", h.getAllList)
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

func (h *GinHandler) get(c *gin.Context) {
	fmt.Println("in")

	c.JSON(201, nil)
}

func (h *GinHandler) getAllList(c *gin.Context) {
	fmt.Println("in")
	c.JSON(201, nil)
}
