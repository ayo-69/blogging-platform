package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCommentsRoutes(r *gin.Engine) {
	commentRoutes := r.Group("/comments")
	{
		commentRoutes.POST("/", controllers.CreateComment)
		commentRoutes.GET("/:id", controllers.GetCommentById)
		commentRoutes.GET("/post/:post_id", controllers.GetCommentsByPostId)
		commentRoutes.GET("/user/:user_id", controllers.GetCommentsByUserId)
		commentRoutes.PUT("/:id", controllers.UpdateCommentById)
		commentRoutes.DELETE("/:id", controllers.DeleteCommentById)
	}
}
