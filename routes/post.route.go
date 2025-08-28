package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	postRoutes := r.Group("/post")
	{
		postRoutes.POST("/", controllers.AddPost)
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPostByID)
		postRoutes.GET("/user/:id", controllers.GetPostsByUserID)
		postRoutes.GET("/tag/:tag", controllers.AddTag)
		postRoutes.PUT("/:id", controllers.UpdatePostByID)
		postRoutes.DELETE("/:id", controllers.DeletePostByID)
	}
}
