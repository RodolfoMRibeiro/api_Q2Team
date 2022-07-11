package controller

import (
	"libary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category

	data, _ := c.GetRawData()

	category.SetCategory(string(data))

	model.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func FindCategory(c *gin.Context) {

	var category model.Category
	var books []model.Book
	var bookshelf []model.Bookshelf

	if err := model.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	bookshelf = findBookshelfByCategoryId(c.Param("id"))

	for _, bookshelfPart := range bookshelf {
		book := model.Book{}
		book.FindById(bookshelfPart.Book)

		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{"category": category, "books": books})
}

func UpdateCategory(c *gin.Context) {

	var category model.Category
	var newCategory model.Category
	jsonElement, _ := c.GetRawData()

	if err := model.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	newCategory.SetCategory(string(jsonElement))

	if model.DB.Model(&category).Updates(&newCategory).RowsAffected == 0 {
		model.DB.Create(&category)
	}
}

func DeleteCategory(c *gin.Context) {
	var category model.Category

	if err := model.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
