package routes

import (
	"golang-comment/controllers"
	"golang-comment/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	comment := r.Group("/api")
	comment.Use(middlewares.JWTAuth())
	comment.GET("/comments", controllers.GetComments)
	comment.POST("/comments", controllers.CreateComment)
	comment.PUT("/comments/:id", controllers.UpdateComment)
	comment.DELETE("/comments/:id", controllers.DeleteComment)
}
