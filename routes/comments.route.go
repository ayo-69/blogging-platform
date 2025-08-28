package routes

import (
	"github.com/ayo-69/blogging-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCommentsRoutes(r *gin.Engine) {
	r.POST("/comments", controllers.CreateComment)
	r.GET("/comments/:id", controllers.GetCommentById)
	r.GET("/comments/post/:post_id", controllers.GetCommentsByPostId)
	r.GET("/comments/user/:user_id", controllers.GetCommentsByUserId)
	r.PUT("/comments/:id", controllers.UpdateCommentById)
	r.DELETE("/comments/:id", controllers.DeleteCommentById)
}
