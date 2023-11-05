package controllers

import (
	"Karchu/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)
	code, msg := user.CreateUser()
	if code == "SUCCESS" {
		ctx.JSON(200, createSuccessResponse(code, msg))
	} else if code == "INVALID_USER_DETAILS" {
		ctx.JSON(400, createErrorResponse(code, msg))
	} else if code == "EMAIL_ALREADY_EXISTS" {
		ctx.JSON(409, createErrorResponse(code, msg))
	} else { //code is "DB_CONNECTIVITY_ISSUE" or "DB_INSERT_ERROR"  or anything {
		ctx.JSON(500, createErrorResponse(code, msg))
	}
}

func GetCategories(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Food":         "",
		"Travel":       "",
		"healthy food": "",
		"Office food":  "",
	})
}

func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}
