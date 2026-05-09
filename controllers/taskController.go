package controllers

import (
	"backend/config"
	"backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {

	var task models.Task

	c.BindJSON(&task)

	userID := c.GetString("user_id")

	objID, _ := primitive.ObjectIDFromHex(userID)

	task.UserID = objID

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, task)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Task creation failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created",
	})
}

func GetTasks(c *gin.Context) {

	userID := c.GetString("user_id")

	objID, _ := primitive.ObjectIDFromHex(userID)

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{
		"user_id": objID,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching tasks",
		})

		return
	}

	var tasks []models.Task

	cursor.All(ctx, &tasks)

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)

	var task models.Task

	c.BindJSON(&task)

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"title":       task.Title,
				"description": task.Description,
				"status":      task.Status,
			},
		},
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Update failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Updated",
	})
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{
		"_id": objID,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Deleted",
	})
}
