package model

type TagCreateInput struct {
	TaskID int    `json:"task_id"`
	Name   string `json:"name"`
}

type TagRemoveInput struct {
	TagId int `json:"tag_id"`
}
