package controller

import (
	model "libary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

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

	if err := model.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBind(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

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
