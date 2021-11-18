package helper

import (
	"Jumia_todoList/api/model"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Unmarshal(c *gin.Context, t interface{}) error {
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(x, &t)

	if err != nil {
		return model.ErrInvalidInput
	}
	return nil
}

func ErrHandler(err error, c *gin.Context) {
	switch err {
	case model.ErrInvalidInput:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}
}
