package model

import "time"

type TaskOutput struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Due         time.Time   `json:"due"`
	Completed   bool        `json:"Completed"`
	Tags        []TagOutput `json:"tags"`
}

type TaskCreateInput struct {
	ListID      uint     `json:"list_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Due         string   `json:"due"`
	Completed   bool     `json:"completed"`
	Tags        []string `json:"tags"`
}

type TaskID struct {
	ID int `json:"id"`
}

type TaskCreateOutput struct {
	Message string `json:"message"`
	Data    TaskID `json:"data"`
}

type TaskUpdateInput struct {
	TaskID      uint   `json:"task_id"`
	ListID      uint   `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Completed   bool   `json:"completed"`
}

type TaskUpdateOutput struct {
	Message string     `json:"message"`
	Data    ListOutput `json:"data"`
}

type TaskRemoveInput struct {
	ID int `json:"id"`
}

type TaskFilterInput struct {
	ListId int      `json:"list_id"`
	Tags   []string `json:"tags"`
}

type TaskFilterOutput struct {
	Data []TaskOutput `json:"data"`
}

type TaskFilterAllInput struct {
	Tags []string `json:"tags"`
}

type TaskFilterAllOutput struct {
	Data []TaskOutput `json:"data"`
}

type GetListTask struct {
	Data []TaskOutput `json:"data"`
}
