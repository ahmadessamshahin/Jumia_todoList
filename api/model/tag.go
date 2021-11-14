package model

type TagOutput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagCreateInput struct {
	TaskID int    `json:"task_id"`
	Name   string `json:"name"`
}
type TagID struct {
	ID int `json:"id"`
}
type TagCreateOutput struct {
	Message string `json:"message"`
	Data    TagID  `json:"data"`
}

type TagRemoveInput struct {
	TagId int `json:"tag_id"`
}
