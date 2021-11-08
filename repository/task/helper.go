package task

import (
	"Jumia_todoList/entity"
	"fmt"
	"strconv"
	"strings"
)

type groupedTags struct {
	TagIds   string `json:"tag_ids"`
	TagNames string `json:"tag_names"`
}

func mapTags(gpTags groupedTags) []entity.Tag {
	tags := make([]entity.Tag, 0)
	ids := strings.Split(gpTags.TagIds[1:len(gpTags.TagIds)-1], ",")
	names := strings.Split(gpTags.TagNames[1:len(gpTags.TagNames)-1], ",")
	for i, v := range ids {
		id, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		tags = append(tags, entity.Tag{Id: int64(id), Name: names[i]})
	}
	return tags
}
