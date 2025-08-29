package controllers

import (
	"net/http"

	"github.com/ayo-69/blogging-platform/database"
	"github.com/ayo-69/blogging-platform/models"
	"github.com/gin-gonic/gin"
)

func GetProfileByID(c *gin.Context) {
	var profile models.Profile

	if err := database.DB.First(&profile, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func GetProfileByUserID(c *gin.Context) {
	var profile models.Profile

	if err := database.DB.First(&profile, "user_id = ?", c.Param("user_id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateProfileByID(c *gin.Context) {
	var profile models.Profile

	if err := database.DB.First(&profile, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func DeleteProfileByID(c *gin.Context) {
	var profile models.Profile

	if err := database.DB.First(&profile, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}
