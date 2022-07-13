package business

import (
	"fmt"
	"libary/model"

	"github.com/gin-gonic/gin"
)

func FindCategory(c *gin.Context) []model.Book {
	var books []model.Book
	var bookshelf []model.Bookshelf

	bookshelf = findBookshelfByCategoryId(c.Param("id"))

	for _, bookshelfPart := range bookshelf {
		book := model.Book{}
		book.FindById(bookshelfPart.Book)
		fmt.Println(bookshelfPart)

		books = append(books, book)
	}

	return books
}
