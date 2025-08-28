package routes

import "github.com/gin-gonic/gin"

func RegisterProfileRoutes(r *gin.Engine) {
	profileRoues := r.Group("/user")
	{
		profileRoues.GET("/:user_id")
	}
}
