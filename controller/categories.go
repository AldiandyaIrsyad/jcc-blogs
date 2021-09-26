package controller

import (
	"log"
	"net/http"

	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	models.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func CreateCategory(c *gin.Context) {
	var input createLabelInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Title: input.Title, Description: input.Description}
	models.DB.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})

}
func UpdateCategory(c *gin.Context) {
	// Get model if exist
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("categoryID")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input createLabelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Line 46")

	models.DB.Model(&category).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	// Get model if exist
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("categoryID")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
