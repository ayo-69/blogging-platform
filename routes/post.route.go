package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	middlware "github.com/ayo-69/blogging-platform/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	postRoutes := r.Group("/post")
	postRoutes.Use(middlware.AuthMiddleware())
	{
		postRoutes.POST("/", controllers.AddPost)
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPostByID)
		postRoutes.GET("/user/:id", controllers.GetPostsByUserID)
		postRoutes.PUT("/:id", controllers.UpdatePostByID)
		postRoutes.DELETE("/:id", controllers.DeletePostByID)
	}
}
