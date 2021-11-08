package model

type ListOutput struct {
	ID   int    `json:"id" copier:"must"`
	Name string `json:"name" copier:"must"`
}

type ListCreateInput struct {
	Name string `json:"name"`
}

type ListCreateOutput struct {
	Message string     `json:"message"`
	Data    ListOutput `json:"data"`
}

type ListUpdateInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListUpdateOutput struct {
	Message string     `json:"message"`
	Data    ListOutput `json:"data"`
}

type ListRemoveInput struct {
	ID int `json:"id"`
}

type GETListOutput struct {
	Data ListOutput `json:"data"`
}

type GetAllListInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetAllListOutput struct {
	Message string       `json:"message"`
	Data    []ListOutput `json:"data"`
}
