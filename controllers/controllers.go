package controllers

import (
	"Karchu/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)
	code, msg := user.AuthenticateUser()
	if code == "AUTHENTICATED" {
		var category models.Category
		ctx.Bind(&category)
		createCode, createMsg := category.CreateCategory(&user)
		if createCode == "SUCCESS" {
			ctx.JSON(200, createSuccessResponse(createCode, createMsg))
		} else if createCode == "CATEGORY_EXISTS" {
			ctx.JSON(401, createErrorResponse(createCode, createMsg))
		} else { //CreateCode is "DB_CONNECTIVITY_ISSUE" or "DB_INSERT_ERROR"  or anything
			ctx.JSON(500, createErrorResponse(createCode, createMsg))
		}
	} else if code == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, createErrorResponse(code, msg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, createErrorResponse(code, msg))
	}
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)
	code, msg := user.CreateUser()
	if code == "SUCCESS" {
		ctx.JSON(200, createSuccessResponse(code, msg))
	} else if code == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, createErrorResponse(code, msg))
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
