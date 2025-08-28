package controllers

import (
	"net/http"

	"github.com/ayo-69/blogging-platform/database"
	"github.com/ayo-69/blogging-platform/models"
	"github.com/gin-gonic/gin"
)

type post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddPost(c *gin.Context) {
	userInput := post{}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("userID")

	// Create a new post
	newPost := models.Post{
		Title:   userInput.Title,
		Content: userInput.Content,
		UserID:  userID,
	}

	// Save the post to the database
	if err := database.DB.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created post
	c.JSON(http.StatusCreated, newPost)

}
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	var post models.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetPostsByUserID(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Where("user_id = ?", c.Param("id")).Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func UpdatePostByID(c *gin.Context) {
	userID := c.GetString("user_id")
	var updatePost models.Post
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&updatePost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userInput := post{}
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePost.Title = userInput.Title
	updatePost.Content = userInput.Content

	if err := database.DB.Save(&updatePost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatePost)
}
func DeletePostByID(c *gin.Context) {
	userID := c.GetString("user_id")
	var post models.Post
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
