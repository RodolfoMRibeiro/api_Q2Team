package controller

import (
	"libary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookshelf(c *gin.Context) {
	var bookshelf model.Bookshelf

	if err := c.ShouldBind(&bookshelf); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	model.DB.Create(&bookshelf)

	c.JSON(http.StatusOK, gin.H{"data": bookshelf})
}

func FindBookshelf(c *gin.Context) { // Get model if exist
	bookshelf := &model.Bookshelf{}

	if err := model.DB.Where("category = ? AND book = ?", c.Param("categoryID"), c.Param("bookID")).First(bookshelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookshelf})
}

func UpdateBookshelf(c *gin.Context) {

	var bookshelf model.Bookshelf
	var newBookshelf model.Bookshelf

	if err := model.DB.Where("id = ?", c.Param("id")).First(&bookshelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBind(&newBookshelf); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	if model.DB.Model(&bookshelf).Updates(&newBookshelf).RowsAffected == 0 {
		model.DB.Create(&bookshelf)
	}

	c.JSON(http.StatusOK, gin.H{"data": bookshelf})
}

func DeleteBookshelf(c *gin.Context) {
	// Get model if exist
	var bookshelf model.Bookshelf
	if err := model.DB.Where("id = ?", c.Param("id")).First(&bookshelf).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	model.DB.Delete(&bookshelf)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
