package controller

import (
	"libary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books

func CreateBook(c *gin.Context) {
	var book model.Book

	data, _ := c.GetRawData()

	book.SetBook(string(data))

	model.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book model.Book

	if err := model.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {

	var book model.Book
	var newBook model.Book
	jsonElement, _ := c.GetRawData()

	if err := model.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	newBook.SetBook(string(jsonElement))

	if model.DB.Model(&book).Updates(&newBook).RowsAffected == 0 {
		model.DB.Create(&book)
	}

}

func DeleteBook(c *gin.Context) {

	var book model.Book
	if err := model.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
