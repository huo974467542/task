package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task4/db"
	"task4/model"
)

func GetCommentsByPostID(c *gin.Context) {
	ID := c.Param("postID")
	postID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var comments []model.Comment
	err = db.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get comments failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func CreateComment(c *gin.Context) {
	ID := c.Param("postID")
	postID, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var comment model.Comment
	err = c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("id")
	comment.PostID = uint(postID)
	comment.UserID = userID
	err = db.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}
