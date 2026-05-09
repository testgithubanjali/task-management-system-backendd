package routes

import (
	"task-management-system-backend/controllers"
	"task-management-system-backend/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {

	task := router.Group("/api/v1/tasks")

	task.Use(middleware.AuthMiddleware())

	{
		task.POST("/", controllers.CreateTask)

		task.GET("/", controllers.GetTasks)

		task.PUT("/:id", controllers.UpdateTask)

		task.DELETE("/:id", controllers.DeleteTask)
	}
}
