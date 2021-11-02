package list

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/usecase/list"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

func (h *GinHandler) get(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}
	i, err := strconv.Atoi(id)

	if err != nil {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}

	list, err := h.UseCase.Get(i)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	var o model.ListOutput
	err = helper.Cast(list, &o)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.GETListOutput{Data: o}
	c.JSON(200, res)
}

//TODO
func (h *GinHandler) getAllList(c *gin.Context) {
	fmt.Println("in")
	c.JSON(201, nil)
}
