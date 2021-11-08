package pkg

import (
	"Jumia_todoList/api/model"
	"bytes"
	"encoding/json"
)

func Cast(x interface{}, t interface{}) error {

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(x)

	err := json.Unmarshal(buffer.Bytes(), &t)

	if err != nil {
		return model.ErrInvalidInput
	}
	return nil
}
