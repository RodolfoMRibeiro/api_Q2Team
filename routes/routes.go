package routes

import (
	"libary/controller"

	"github.com/gin-gonic/gin"
)

func Groups(router *gin.Engine) {
	BookGroup := router.Group("/book")
	{
		BookGroup.POST("/", controller.CreateBook)
		BookGroup.GET("/:id", controller.FindBook)
		BookGroup.PUT("/:id", controller.UpdateBook)
		BookGroup.DELETE("/:id", controller.DeleteBook)
	}

	CategoryGroup := router.Group("/category")
	{
		CategoryGroup.POST("/", controller.CreateCategory)
		CategoryGroup.GET("/:id", controller.FindCategory)
		CategoryGroup.PUT("/:id", controller.UpdateCategory)
		CategoryGroup.DELETE("/:id", controller.DeleteCategory)
	}

	BookshelfGroup := router.Group("/bookshelf")
	{
		BookshelfGroup.POST("/", controller.CreateBookshelf)
		BookshelfGroup.GET("/:categoryID/:bookID", controller.FindBookshelf)
		BookshelfGroup.PUT("/:id", controller.UpdateBookshelf)
		BookshelfGroup.DELETE("/:id", controller.DeleteBookshelf)
	}
}
