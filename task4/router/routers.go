package router

import (
	"github.com/gin-gonic/gin"
	"task4/controller"
	"task4/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/posts", controller.GetAllPosts)
	r.GET("/post/:id", controller.GetPostByID)
	r.GET("/post/comments/:postID", controller.GetCommentsByPostID)
	r.Use(middleware.JWTAuth)
	r.POST("/post", controller.CreatePost)
	r.DELETE("/post/:id", controller.DeletePostByID)
	r.PUT("/post/:id", controller.UpdatePost)
	r.POST("/comment/:postID", controller.CreateComment)
}
