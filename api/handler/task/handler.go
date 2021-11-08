package task

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"Jumia_todoList/usecase/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinHandler struct {
	UseCase task.UseCase
}

func (h *GinHandler) Routes(router *gin.Engine, useCase interface{}) {
	h.UseCase = useCase.(task.UseCase)

	taskRouter := router.Group("/task")
	{
		taskRouter.GET("/", h.get)
		taskRouter.POST("/", h.add)
		taskRouter.PATCH("/", h.edit)
		taskRouter.DELETE("/", h.remove)
		taskRouter.GET("/filter", h.filter)
	}
}

func (h *GinHandler) add(c *gin.Context) {
	var l model.TaskCreateInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Task %s created successfully", l.Title)}
	c.JSON(204, res)
}

func (h *GinHandler) edit(c *gin.Context) {
	var l model.TaskUpdateInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Task %s updated successfully", l.Title)}
	c.JSON(204, res)
}

func (h *GinHandler) remove(c *gin.Context) {
	var l model.TaskRemoveInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Task %s updated successfully", l.ID)}
	c.JSON(204, res)

}

func (h *GinHandler) filter(c *gin.Context) {
	var l model.TaskFilterInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	tasks, err := h.UseCase.Filter(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.TaskFilterOutput{Data: tasks}
	c.JSON(200, res)

}

func (h *GinHandler) get(c *gin.Context) {
	id, ok := c.GetQuery("id")
	fmt.Println(id)
	if !ok {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}
	i, err := strconv.Atoi(id)

	if err != nil {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return
	}

	tasks := h.UseCase.Get(i)

	res := model.GETListOutput{Data: tasks}
	c.JSON(200, res)
}
