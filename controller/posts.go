package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post

	var comments []models.Comment

	if err := models.DB.Where("id = ?", c.Param("postID")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Where("post_id = ?", c.Param("postID")).Find(&comments)

	post.Comments = append(post.Comments, comments...)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UserCreatePost(c *gin.Context) {
	var input createPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID_INT, _ := strconv.Atoi(c.Param("userID"))
	userID := uint(userID_INT)
	// postID_INT, _ := strconv.Atoi(c.Param("postID"))
	// postID := uint(postID_INT)

	// Categories
	var categories []models.Category

	log.Println(input.Categories)
	log.Println(input.Tags)
	for _, category_name := range input.Categories {
		var category models.Category
		models.DB.Where("title = ?", category_name).First(&category)

		if category.Title == category_name {
			categories = append(categories, category)

		}
	}

	var tags []models.Tag
	for _, tag_name := range input.Tags {
		var tag models.Tag
		models.DB.Where("title = ?", tag_name).First(&tag)
		if tag.Title == tag_name {
			tags = append(tags, tag)
		}
	}

	post := models.Post{Title: input.Title, UserID: userID, Summary: input.Summary, Content: input.Content, Categories: categories, Tags: tags}
	models.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})

}
func UpdatePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("postID")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input createPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&post).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("postID")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
