package controller

import (
	"net/http"

	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context) {
	var tags []models.Tag

	models.DB.Find(&tags)
	c.JSON(http.StatusOK, gin.H{"data": tags})
}

func CreateTag(c *gin.Context) {
	var input createLabelInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{Title: input.Title, Description: input.Description}
	models.DB.Create(&tag)
	c.JSON(http.StatusOK, gin.H{"data": tag})

}
func UpdateTag(c *gin.Context) {
	// Get model if exist
	var tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("categoryID")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input createPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&tag).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

func DeleteTag(c *gin.Context) {
	// Get model if exist
	var tag models.Category
	if err := models.DB.Where("id = ?", c.Param("categoryID")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
