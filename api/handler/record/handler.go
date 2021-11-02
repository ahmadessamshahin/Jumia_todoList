package record

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
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
		recordRouter.GET("/filter", h.filter)
	}
}

func (h *GinHandler) add(c *gin.Context) {
	var l model.RecordCreateInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Record %s created successfully", l.Title)}
	c.JSON(204, res)
}

func (h *GinHandler) edit(c *gin.Context) {
	var l model.RecordUpdateInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Record %s updated successfully", l.Title)}
	c.JSON(204, res)
}

func (h *GinHandler) remove(c *gin.Context) {
	var l model.RecordRemoveInput
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

	res := model.EmptySuccessfulOutput{Message: fmt.Sprintf("Record %s updated successfully", l.ID)}
	c.JSON(204, res)

}

func (h *GinHandler) filter(c *gin.Context) {
	var l model.RecordFilterInput
	err := helper.Unmarshal(c, &l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	records, err := h.UseCase.Filter(l)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}
	var o []model.RecordOutput
	err = helper.Cast(records, &o)

	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	res := model.RecordFilterOutput{Data: o}
	c.JSON(200, res)

}
