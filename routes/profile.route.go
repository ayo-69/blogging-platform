package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine) {
	profileRoutes := r.Group("/profile")
	{
		profileRoutes.GET("/:id", controllers.GetProfileByID)
		profileRoutes.GET("/user/:id", controllers.GetProfileByUserID)
		profileRoutes.PUT("/:id", controllers.UpdateProfileByID)
		profileRoutes.DELETE("/:id", controllers.DeleteProfileByID)
	}
}
