package model

type TagOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TagCreateInput struct {
	Tags []string `json:"tags"`
}

type TagCreateOutput struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"code"`
	Data       []TagOutput `json:"data"`
}

type TagUpdateInput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TagUpdateOutput struct {
	Message    string    `json:"message"`
	StatusCode int       `json:"code"`
	Data       TagOutput `json:"data"`
}

type TagDeleteInput int

type TagDeleteOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}
