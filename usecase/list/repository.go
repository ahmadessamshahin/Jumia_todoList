package list

import "Jumia_todoList/entity"

type ListingRepository interface {
	CreateList() entity.List
}
