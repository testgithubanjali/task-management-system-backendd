package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {

	admin := router.Group("/api/v1/admin")

	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminOnly())

	{
		admin.GET("/users", controllers.GetAllUsers)

		admin.DELETE("/users/:id", controllers.DeleteUser)
	}
}
