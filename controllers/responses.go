package controllers

import "github.com/gin-gonic/gin"

func createErrorResponse(errorCode string, errorMessage string) gin.H {
	return (gin.H{
		"error_code":    errorCode,
		"error_message": errorMessage,
	})
}

func createSuccessResponse(successCode string, successMessage string) gin.H {
	return (gin.H{
		"success_code":    successCode,
		"success_message": successMessage,
	})
}
