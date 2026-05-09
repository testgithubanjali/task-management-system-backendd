package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/api/v1/auth")

	{
		auth.POST("/register", controllers.Register)

		auth.POST("/login", controllers.Login)
	}
}
