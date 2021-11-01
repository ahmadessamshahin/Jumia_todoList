package model

type ListOutput struct {
	Name    string         `json:"name"`
	Records []RecordOutput `json:"records"`
}

type ListCreateInput struct {
	Name string `json:"name"`
}

type ListCreateOutput struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"code"`
	Data       ListOutput `json:"data"`
}

type ListUpdateInput struct {
	Name string `json:"name"`
}

type ListUpdateOutput struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"code"`
	Data       ListOutput `json:"data"`
}

type GetAllListInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetAllListOutput struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"code"`
	Data       []ListOutput `json:"data"`
}
