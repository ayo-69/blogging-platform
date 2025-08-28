package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterTagsRoutes(r *gin.Engine) {
	tagsRoutes := r.Group("/tags")
	{
		tagsRoutes.POST("/", controllers.AddTag)
		tagsRoutes.GET("/", controllers.GetTags)
		tagsRoutes.GET("/:tag_id", controllers.GetTagByID)
	}
}
