package tag

import "Jumia_todoList/entity"

type TaggingRepository interface {
	Create() *entity.Tag
	Delete() error
}
