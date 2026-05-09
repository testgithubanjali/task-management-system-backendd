package controllers

import (
	"context"
	"net/http"
	"task-management-system-backend/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch users",
		})

		return
	}

	var users []bson.M

	cursor.All(ctx, &users)

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid User ID",
		})

		return
	}

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{
		"_id": objID,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete Failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Deleted",
	})
}
