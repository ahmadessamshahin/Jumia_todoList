package list

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/usecase/list"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}

	id, err := h.UseCase.Create(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.ListCreateOutput{Message: "success", Data: model.ListID{ID: id}}
	c.JSON(201, res)
}

func (h *GinHandler) edit(c *gin.Context) {
	var l model.ListUpdateInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}

	err = h.UseCase.Update(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.EmptySuccessfulOutput{Message: "success"}
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

	res := model.EmptySuccessfulOutput{Message: "success"}
	c.JSON(200, res)
}

func (h *GinHandler) getAllList(c *gin.Context) {
	i := validateGetAllInput(c)
	if i == nil {
		return
	}

	lists := h.UseCase.GetAllLists(*i)

	var o []model.ListOutput
	err := copier.Copy(&o, &lists)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.GetAllListOutput{Message: "success", Data: o}

	c.JSON(200, res)
}
