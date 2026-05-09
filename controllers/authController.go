package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Data",
		})

		return
	}

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var existingUser models.User

	err := collection.FindOne(ctx, bson.M{
		"email": user.Email,
	}).Decode(&existingUser)

	if err == nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})

		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)

	user.Password = hashedPassword

	if user.Role == "" {
		user.Role = "user"
	}

	_, err = collection.InsertOne(ctx, user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Registration failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered",
	})
}

func Login(c *gin.Context) {

	var data map[string]string

	c.BindJSON(&data)

	email := data["email"]

	password := data["password"]

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var user models.User

	err := collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Credentials",
		})

		return
	}

	isPasswordCorrect := utils.CheckPassword(password, user.Password)

	if !isPasswordCorrect {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Credentials",
		})

		return
	}

	token, _ := utils.GenerateToken(user.ID.Hex(), user.Role)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
