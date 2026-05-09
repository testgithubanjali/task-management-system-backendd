package controllers

import (
	"backend/config"
	"backend/models"
	"backend/validators"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {

	var input validators.TaskInput

	if err := c.BindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Input",
		})

		return
	}

	if !validators.ValidateStruct(c, input) {
		return
	}

	userID := c.GetString("user_id")

	objID, _ := primitive.ObjectIDFromHex(userID)

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		UserID:      objID,
	}

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.InsertOne(ctx, task)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Task Creation Failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task Created Successfully",
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
			"message": "Failed to Fetch Tasks",
		})

		return
	}

	var tasks []models.Task

	cursor.All(ctx, &tasks)

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Task ID",
		})

		return
	}

	var input validators.TaskInput

	if err := c.BindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Input",
		})

		return
	}

	if !validators.ValidateStruct(c, input) {
		return
	}

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"title":       input.Title,
				"description": input.Description,
				"status":      input.Status,
			},
		},
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Task Update Failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Updated Successfully",
	})
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Task ID",
		})

		return
	}

	collection := config.DB.Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{
		"_id": objID,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Task Delete Failed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task Deleted Successfully",
	})
}
