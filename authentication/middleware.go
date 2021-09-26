package authentication

import (
	"net/http"
	"strconv"

	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

type parameter struct {
	c    *gin.Context
	user models.User
}

// login

func Login(c *gin.Context) {
	username, password, hasAuth := c.Request.BasicAuth()

	if !hasAuth {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{"error": "required to be loggedin"})

		return
	}

	var user models.User
	models.DB.Where("username = ?", username).First(&user)

	if user.Password != password {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong username or password"})

		return
	}

	if c.Param("userID") != "" {
		if c.Param("userID") != strconv.Itoa(int(user.ID)) {
			c.Abort()
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized on other account"})
		}
	}

	user_data := []gin.Param{
		{Key: "UD_username", Value: username},
		{Key: "UD_password", Value: password},
	}
	c.Params = append(c.Params, user_data...)
	c.Next()
}

func auth(c *gin.Context, perms []func(p parameter) bool) {

	var user models.User
	models.DB.Where("username = ?", c.Param("UD_username")).First(&user)

	// check permission

	for _, f := range perms {
		if f(parameter{c, user}) {
			c.Next()
			return
		}
	}

	c.Abort()
	c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	c.JSON(http.StatusBadRequest, gin.H{"error": "auth_unauthorized"})

}

func isAdmin(p parameter) bool     { return p.user.Privillege <= 1 && p.user.Privillege != 0 }
func isPoster(p parameter) bool    { return p.user.Privillege <= 2 && p.user.Privillege != 0 }
func isCommenter(p parameter) bool { return p.user.Privillege <= 3 && p.user.Privillege != 0 }

func IsAdmin(c *gin.Context) {
	perms := []func(p parameter) bool{
		isAdmin,
	}

	auth(c, perms)
}

func isPostOP(p parameter) bool {

	var post models.Post
	if err := models.DB.Where("id = ?", p.c.Param("postID")).First(&post).Error; err != nil {
		return false
	}

	return p.user.ID == post.UserID
}
func HandlerCreatePost(c *gin.Context) {
	perms := []func(p parameter) bool{
		isPoster,
	}

	auth(c, perms)
}

func HandlerEditPost(c *gin.Context) {
	perms := []func(p parameter) bool{
		isPostOP,
	}

	auth(c, perms)
}

func HandlerDeletePost(c *gin.Context) {
	perms := []func(p parameter) bool{
		isPostOP,
		isAdmin,
	}

	auth(c, perms)
}

func isCommentOP(p parameter) bool {

	var comment models.Comment
	if err := models.DB.Where("id = ?", p.c.Param("commentID")).First(&comment).Error; err != nil {
		return false
	}

	return p.user.ID == comment.UserID
}

func HandlerCreateComment(c *gin.Context) {
	perms := []func(p parameter) bool{
		isCommenter,
	}

	auth(c, perms)
}

func HandlerDeleteComment(c *gin.Context) {
	perms := []func(p parameter) bool{
		isCommentOP,
		isPostOP,
		isAdmin,
	}

	auth(c, perms)
}
