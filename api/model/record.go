package model

import "time"

type RecordOutput struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Due         time.Time   `json:"due"`
	Completed   bool        `json:"Completed"`
	Tags        []TagOutput `json:"tags"`
}

type RecordCreateInput struct {
	ListID      string `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Completed   bool   `json:"completed"`
}

type RecordUpdateInput struct {
	ListID      string `json:"list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Completed   bool   `json:"completed"`
}

type RecordUpdateOutput struct {
	Message string     `json:"message"`
	Data    ListOutput `json:"data"`
}

type RecordRemoveInput struct {
	ID string `json:"id"`
}

type RecordFilterInput struct {
	ListId string   `json:"list_id"`
	Tags   []string `json:"tags"`
}

type RecordFilterOutput struct {
	Data []RecordOutput `json:"data"`
}
