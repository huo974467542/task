package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task4/db"
	"task4/model"
)

func CreatePost(c *gin.Context) {
	var request struct {
		Title   string
		Content string
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"request": c.Request})
		return
	}
	post := model.Post{
		Title:   request.Title,
		Content: request.Content,
		UserID:  c.GetUint("id"),
	}
	err = db.DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create post failed"})
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func GetAllPosts(c *gin.Context) {
	var posts []model.Post
	err := db.DB.Preload("User").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get all posts failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	err := db.DB.Preload("User").Find(&post, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get post by id failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func DeletePostByID(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	err := db.DB.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get post by id failed"})
		return
	}
	userID := c.GetUint("id")
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "don't have the right to delete"})
		return
	}
	err = db.DB.Delete(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete post by id failed"})
		return
	}
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	err := db.DB.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Get post by id failed"})
		return
	}
	userID := c.GetUint("id")
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "don't have the right to update"})
		return
	}
	var req model.Post
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.Title = req.Title
	post.Content = req.Content
	err = db.DB.Save(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update post by id failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}
