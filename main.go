package main

import (
	"log"
	"os"
	"task-management-system-backend/config"
	"task-management-system-backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB(os.Getenv("MONGO_URI"))

	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{

		AllowOrigins: []string{
			"http://localhost:5174",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	// Routes
	routes.AuthRoutes(router)

	routes.TaskRoutes(router)

	routes.AdminRoutes(router)

	// Health Check Route
	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Backend Running",
		})
	})

	// Start Server
	router.Run(":" + os.Getenv("PORT"))
}
