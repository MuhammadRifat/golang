package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ValidationErr map[string]string
}

type ValidationErrorStruct struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ValidationErr map[string]string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFoundErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusNotFound, Message: msg[0]}
	}
	return &AppError{Code: http.StatusNotFound, Message: "Not found"}
}

func BadRequestErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusBadRequest, Message: msg[0]}
	}
	return &AppError{Code: http.StatusBadRequest, Message: "Bad request"}
}

func UnauthorizedErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusUnauthorized, Message: msg[0]}
	}
	return &AppError{Code: http.StatusBadRequest, Message: "Unauthorized"}
}

func InternalServerErr(msg ...string) *AppError {
	if len(msg) > 0 {
		return &AppError{Code: http.StatusInternalServerError, Message: msg[0]}
	}
	return &AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
}

func ValidationErr(err error) *AppError {
	errorMessages := make(map[string]string)

	// Type assertion to get validation errors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := fieldErr.Field() // Get the struct field name
			tag := fieldErr.Tag()     // Get the validation tag (e.g., "required", "url")
			switch tag {
			case "required":
				errorMessages[field] = "This field is required."
			case "url":
				errorMessages[field] = "Invalid URL format."
			case "email":
				errorMessages[field] = "Invalid email format."
			default:
				errorMessages[field] = "Invalid value."
			}
		}
	}

	return &AppError{Code: http.StatusBadRequest, Message: "Bad request", ValidationErr: errorMessages}
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			switch e := err.(type) {
			case *AppError:
				c.AbortWithStatusJSON(e.Code, gin.H{
					"StatusCode": e.Code,
					"Error":      e.Message,
					"Messages":   e.ValidationErr,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"StatusCode": http.StatusInternalServerError,
					"Error":      e.Error(),
				})
			}
		}
	}
}
