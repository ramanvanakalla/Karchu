package controllers

import (
	"Karchu/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var categoryEntry struct {
		Email        string
		Password     string
		CategoryName string
	}
	ctx.Bind(&categoryEntry)
	user := models.User{Email: categoryEntry.Email, Password: categoryEntry.Password}
	category := models.Category{CategoryName: categoryEntry.CategoryName}
	code, msg := user.AuthenticateUser()
	if code == "AUTHENTICATED" {
		createCode, createMsg := category.CreateCategory(&user)
		if createCode == "SUCCESS" {
			ctx.JSON(200, createSuccessResponse(createCode, createMsg))
		} else if createCode == "CATEGORY_EXISTS" || createCode == "INVALID_CATEGORY" {
			ctx.JSON(401, createErrorResponse(createCode, createMsg))
		} else { //CreateCode is "DB_CONNECTIVITY_ISSUE" or "DB_INSERT_ERROR"  or anything
			ctx.JSON(500, createErrorResponse(createCode, createMsg))
		}
		return
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
	var user models.User
	ctx.Bind(&user)
	authCode, authMsg := user.AuthenticateUser()
	if authCode == "AUTHENTICATED" {
		if categoriesArr, err := user.GetCategories(); err != nil {
			ctx.JSON(500, createErrorResponse("DB_CONNECTIVITY_ISSUE", err.Error()))
		} else {
			ctx.JSON(200, categoriesArr)
		}

	} else if authCode == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, createErrorResponse(authCode, authMsg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, createErrorResponse(authCode, authMsg))
	}

}

func NewTransaction(ctx *gin.Context) {
	var transactionEntry struct {
		models.User
		models.Transaction
	}
	ctx.Bind(&transactionEntry)
	user := models.User{Email: transactionEntry.Email, Password: transactionEntry.Password}
	authCode, authMsg := user.AuthenticateUser()
	if authCode == "AUTHENTICATED" {
		transaction := models.Transaction{UserId: user.ID, Time: transactionEntry.Time, Amount: transactionEntry.Amount, Category: transactionEntry.Category, Description: transactionEntry.Description, SplitTag: transactionEntry.SplitTag, MapUrl: transactionEntry.MapUrl}
		if msg, err := transaction.NewTransaction(); err != nil {
			ctx.JSON(500, createErrorResponse("INSERT", err.Error()))
		} else {
			ctx.JSON(200, createSuccessResponse("SUCCESS", msg))
		}
	} else if authCode == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, createErrorResponse(authCode, authMsg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, createErrorResponse(authCode, authMsg))
	}
}

func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}
