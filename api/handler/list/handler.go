package list

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/usecase/list"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase list.UseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {

	h.UseCase = useCase.(list.UseCase)

	listRouter := router.Group("/list")
	{
		listRouter.POST("/", h.add)
		listRouter.PATCH("/", h.edit)
		listRouter.DELETE("/", h.remove)
		listRouter.GET("/all", h.getAllList)
	}
}

func (h *GinHandler) add(c *gin.Context) {
	var l model.ListCreateInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	err = h.UseCase.Create(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("List %s created successfully", l.Name)}
	c.JSON(204, res)
}

func (h *GinHandler) edit(c *gin.Context) {
	var l model.ListUpdateInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	err = h.UseCase.Update(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("List %s updated successfully", l.Name)}
	c.JSON(200, res)
}

func (h *GinHandler) remove(c *gin.Context) {
	var l model.ListRemoveInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("List %d deleted successfully", l.ID)}
	c.JSON(200, res)
}

func (h *GinHandler) getAllList(c *gin.Context) {
	i := validateGetAllInput(c)
	if i == nil {
		return
	}

	lists := h.UseCase.GetAllLists(*i)
	fmt.Println(lists)

	res := model.GetAllListOutput{Message: fmt.Sprintf("All list limit: %d, Offset: %d", i.Limit, i.Offset),
		Data: lists}

	c.JSON(201, res)
}
