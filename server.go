package main

import (
	"log"

	"github.com/Dragonopy/jcc-blogs/authentication"
	"github.com/Dragonopy/jcc-blogs/controller"
	"github.com/Dragonopy/jcc-blogs/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	models.ConnectDB()

	// Setup DB
	models.DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Category{}, &models.Tag{}, &models.Comment{})

	// create admin
	{
		var admin models.User
		if err := models.DB.Where("username = ?", "aldiandya").First(&admin).Error; err == nil {
			log.Println("admin already exist")
		} else {
			user := models.User{Username: "aldiandya", Password: "password", Privillege: 1}
			models.DB.Create(&user)

		}
	}

	// get/saved on users
	server.GET("/users", authentication.Login, authentication.IsAdmin, controller.GetAllUsers)
	server.POST("/users", controller.CreateUser)
	server.PUT("/users/:userID/changePassword", authentication.Login, controller.ChangePassword)
	server.PUT("/edituser/:id", authentication.Login, authentication.IsAdmin, controller.UpdateUser)
	server.DELETE("/deleteuser/:id", authentication.Login, authentication.IsAdmin, controller.DeletePost)

	// get/saved on posts
	server.GET("/posts", authentication.Login, controller.GetAllPosts)
	server.GET("/posts/:postID", authentication.Login, controller.GetPost)
	server.POST("/users/:userID/posts", authentication.Login, authentication.HandlerCreatePost, controller.UserCreatePost)
	server.PUT("/users/:userID/posts/:postID", authentication.Login, authentication.HandlerEditPost, controller.UpdatePost)
	server.DELETE("/users/:userID/posts/:postID", authentication.Login, authentication.HandlerDeletePost, controller.DeletePost)
	server.DELETE("posts/:postID", authentication.Login, authentication.HandlerDeletePost, controller.DeletePost)

	//  get/saved on comments
	server.POST("/users/:userID/posts/:postID/comments", authentication.Login, controller.UserOnPostCreateComments)
	server.DELETE("/users/:userID/posts/:postID/comments/:commentID", authentication.Login, authentication.HandlerDeleteComment, controller.UserOnPostDeleteComment)
	server.DELETE("posts/:postID/comments/:commentID", authentication.Login, authentication.HandlerDeleteComment, controller.UserOnPostDeleteComment)

	// categories and tags are very similar in terms of structure
	// usage is kinda different category is more wide
	// tags are more spesific
	// E.g category: Programming Tags: Golang, gin

	// get/saved on categories
	server.GET("/categories/", authentication.Login, authentication.IsAdmin, controller.GetAllCategories)
	server.POST("/categories", authentication.Login, authentication.IsAdmin, controller.CreateCategory)
	server.PUT("/categories/:categoryID", authentication.Login, authentication.IsAdmin, controller.UpdateCategory)
	server.DELETE("/categories/:categoryID", authentication.Login, authentication.IsAdmin, controller.DeleteCategory)

	// get/saved on tags
	server.GET("/tags/", authentication.Login, controller.GetAllCategories)
	server.POST("/tags", authentication.Login, authentication.IsAdmin, controller.CreateTag)
	server.PUT("/tags/:tagID", authentication.Login, authentication.IsAdmin, controller.UpdateTag)
	server.DELETE("/tags/:tagID", authentication.Login, authentication.IsAdmin, controller.DeleteTag)

	// Tag are label that can be created by poster

	server.Run(":8080")
}
