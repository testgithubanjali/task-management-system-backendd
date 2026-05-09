package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TaskInput struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=5"`
	Status      string `json:"status" validate:"required"`
}

func ValidateStruct(c *gin.Context, data interface{}) bool {

	err := Validate.Struct(data)

	if err != nil {

		validationErrors := err.(validator.ValidationErrors)

		errors := make(map[string]string)

		for _, fieldError := range validationErrors {

			field := fieldError.Field()

			switch fieldError.Tag() {

			case "required":
				errors[field] = field + " is required"

			case "email":
				errors[field] = "Invalid email format"

			case "min":
				errors[field] = field + " is too short"

			case "max":
				errors[field] = field + " is too long"

			default:
				errors[field] = "Invalid value"
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errors":  errors,
		})

		return false
	}

	return true
}
