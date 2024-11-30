package routes

import (
	"golang-comment/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("auth/register", controllers.Register)
	r.POST("auth/login", controllers.Login)
}
