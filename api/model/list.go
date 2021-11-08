package model

import "Jumia_todoList/entity"

type ListCreateInput struct {
	Name string `json:"name"`
}

type ListCreateOutput struct {
	Message string      `json:"message"`
	Data    entity.List `json:"data"`
}

type ListUpdateInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListUpdateOutput struct {
	Message string      `json:"message"`
	Data    entity.List `json:"data"`
}

type ListRemoveInput struct {
	ID int `json:"id"`
}

type GETListOutput struct {
	Data []entity.Task `json:"data"`
}

type GetAllListInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetAllListOutput struct {
	Message string        `json:"message"`
	Data    []entity.List `json:"data"`
}
