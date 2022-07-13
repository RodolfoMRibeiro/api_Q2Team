package controller

import (
	"libary/business"

	"libary/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	model.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func FindCategory(c *gin.Context) {

	var category model.Category
	var books []model.Book

	if err := model.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	books = business.FindCategory(c)

	c.JSON(http.StatusOK, gin.H{"category": category, "books": books})
}

func UpdateCategory(c *gin.Context) {

	var category model.Category
	var newCategory model.Category

	if err := model.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBind(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

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
