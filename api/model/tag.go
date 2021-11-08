package model

type TagOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagCreateInput struct {
	TaskID int    `json:"task_id"`
	Name   string `json:"name"`
}

type TagRemoveInput struct {
	TagId int `json:"tag_id"`
}
