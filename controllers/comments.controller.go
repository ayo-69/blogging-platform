package controllers

import (
	"github.com/ayo-69/blogging-platform/database"
	"github.com/ayo-69/blogging-platform/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, comment)
}

func GetCommentById(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comment)
}

func GetCommentsByPostId(c *gin.Context) {
	id := c.Param("post_id")
	var comments []models.Comment

	if err := database.DB.Where("post_id = ?", id).Find(&comments).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}

func GetCommentsByUserId(c *gin.Context) {
	id := c.Param("user_id")
	var comments []models.Comment

	if err := database.DB.Where("user_id = ?", id).Find(&comments).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}

func UpdateCommentById(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comment)
}

func DeleteCommentById(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}
