package model

type ErrorOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
