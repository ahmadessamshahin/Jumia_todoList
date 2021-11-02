package model

type TagOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TagCreateInput struct {
	RecordID string   `json:"record_id"`
	Tags     []string `json:"tags"`
}

type TagRemoveInput struct {
	TagId    string `json:"tag_id"`
	RecordId string `json:"record_id"`
}
