package controller

import (
	"net/http"
	"strconv"

	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

func UserOnPostCreateComments(c *gin.Context) {
	var input createCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if post exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("postID")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Post ID"})
		return
	}

	userID_INT, _ := strconv.Atoi(c.Param("userID"))
	userID := uint(userID_INT)

	postID_INT, _ := strconv.Atoi(c.Param("postID"))
	postID := uint(postID_INT)

	comment := models.Comment{UserID: userID, PostID: postID, Content: input.Content}
	models.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{"data": comment})

}

func UserOnPostDeleteComment(c *gin.Context) {
	// Get model if exist
	var comment models.Comment
	if err := models.DB.Where("id = ?", c.Param("commentID")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
