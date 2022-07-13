package business

import (
	"libary/model"
)

func findBookshelfByCategoryId(id string) []model.Bookshelf {
	bookshelf := &[]model.Bookshelf{}

	model.DB.Where("category = ?", id).Find(bookshelf)

	return *bookshelf
}
