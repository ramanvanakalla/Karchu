package responses

import "github.com/gin-gonic/gin"

type ErrorRes struct {
	Error_code    string
	Error_message string
}

type SuccessRes struct {
	Success_code    string
	Success_message string
}

func CreateErrorResponse(errorCode string, errorMessage string) gin.H {
	return (gin.H{
		"error_code":    errorCode,
		"error_message": errorMessage,
	})
}

func CreateSuccessResponse(successCode string, successMessage string) gin.H {
	return (gin.H{
		"success_code":    successCode,
		"success_message": successMessage,
	})
}
