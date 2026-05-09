package controllers

import (
	"context"
	"net/http"
	"task-management-system-backend/config"
	"task-management-system-backend/models"
	"task-management-system-backend/utils"
	"task-management-system-backend/validators"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *gin.Context) {

	var input validators.RegisterInput

	if err := c.BindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Input",
		})

		return
	}

	if !validators.ValidateStruct(c, input) {
		return
	}

	collection := config.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var existingUser models.User

	err := collection.FindOne(ctx, bson.M{
		"email": input.Email,
	}).Decode(&existingUser)

	if err == nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})

		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
	}

	if user.Role == "" {
		user.Role = "user"
	}

	_, err = collection.InsertOne(ctx, user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Registration Failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered Successfully",
	})
}

func Login(c *gin.Context) {

	var input validators.LoginInput

	if err := c.BindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Input",
		})

		return
	}

	if !validators.ValidateStruct(c, input) {
		return
	}

	email := input.Email
	password := input.Password

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
		"message": "Login Successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
