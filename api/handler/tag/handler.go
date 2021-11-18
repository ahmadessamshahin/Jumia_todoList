package tag

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/usecase/tag"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase tag.UseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {
	h.UseCase = useCase.(tag.UseCase)

	tagRouter := router.Group("/tag")
	{
		tagRouter.POST("/", h.add)
		tagRouter.DELETE("/", h.remove)
	}
}

func (h *GinHandler) add(c *gin.Context) {
	var l model.TagCreateInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	_, err = h.UseCase.Create(l)
	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.EmptySuccessfulOutput{Message: "Success"}
	c.JSON(201, res)
}

func (h *GinHandler) remove(c *gin.Context) {
	var l model.TagRemoveInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	err = h.UseCase.Delete(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Tags created successfully")}
	c.JSON(204, res)
}
