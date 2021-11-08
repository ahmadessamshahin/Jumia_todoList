package list

import (
	"Jumia_todoList/api/helper"
	"Jumia_todoList/api/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func validateGetAllInput(c *gin.Context) *model.GetAllListInput {
	limit, ok := c.GetQuery("limit")
	if !ok {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return nil
	}

	offset, ok := c.GetQuery("offset")
	if !ok {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return nil
	}

	limitVal, err := strconv.Atoi(limit)

	if err != nil {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return nil
	}

	offsetVal, err := strconv.Atoi(offset)

	if err != nil {
		helper.ErrHandler(model.ErrInvalidInput, c)
		return nil
	}
	return &model.GetAllListInput{Limit: limitVal, Offset: offsetVal}
}
