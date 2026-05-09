package main

import (
	"log"
	"os"
	"task-management-system-backend/config"
	"task-management-system-backend/routes"

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

	routes.AuthRoutes(router)
	routes.TaskRoutes(router)
	routes.AdminRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
