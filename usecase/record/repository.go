package record

import (
	"Jumia_todoList/entity"
)

type RecordingRepository interface {
	Create() *entity.Record
	Update() *entity.Record
	Delete() error
	FilterList() *entity.Record
	FilterInAllList() *entity.Record
}
