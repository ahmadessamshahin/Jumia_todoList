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
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due         time.Time `json:"due"`
	Completed   bool      `json:"Completed"`
}

type RecordCreateOutput struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"code"`
	Data       RecordOutput `json:"data"`
}

type RecordUpdateInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due         time.Time `json:"due"`
	Completed   bool      `json:"Completed"`
}

type RecordUpdateOutput struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"code"`
	Data       ListOutput `json:"data"`
}

type RecordDeleteInput int

type RecordDeleteOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}
