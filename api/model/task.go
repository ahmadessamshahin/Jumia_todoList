package model

import (
	"Jumia_todoList/entity"
)

type TaskCreateInput struct {
	ListID      uint     `json:"list_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Due         string   `json:"due"`
	Completed   bool     `json:"completed"`
	Tags        []string `json:"tags"`
}

type TaskUpdateInput struct {
	TaskID      uint   `json:"task_id"`
	ListID      uint   `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Completed   bool   `json:"completed"`
}

type TaskRemoveInput struct {
	ID int `json:"id"`
}

type TaskFilterInput struct {
	ListId int      `json:"list_id"`
	Tags   []string `json:"tags"`
}

type TaskFilterOutput struct {
	Data []entity.Task `json:"data"`
}

type TaskFilterAllInput struct {
	Tags []string `json:"tags"`
}

type TaskFilterAllOutput struct {
	Data []entity.Task `json:"data"`
}

type GetListTask struct {
	Data []entity.Task `json:"data"`
}
