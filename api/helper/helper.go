package helper

import (
	"Jumia_todoList/api/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Unmarshal(c *gin.Context, t interface{}) error {
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(x, &t)

	if err != nil {
		return model.ErrInvalidInput
	}
	return nil
}

func Cast(x interface{}, t interface{}) error {

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(x)

	err := json.Unmarshal(buffer.Bytes(), &t)

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
		c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
	}
}
