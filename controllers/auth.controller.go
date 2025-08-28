package controllers

import (
	"net/http"

	"github.com/ayo-69/blogging-platform/database"
	"github.com/ayo-69/blogging-platform/models"
	"github.com/ayo-69/blogging-platform/utils"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userInput user
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username:     userInput.Username,
		Email:        userInput.Email,
		PasswordHash: hashPassword,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user, "token": token})
}

func Login(c *gin.Context) {
	var userInput user
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if err := database.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(userInput.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
