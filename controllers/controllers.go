package controllers

import (
	"Karchu/helpers"
	"Karchu/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateCategory(ctx *gin.Context) {
	var categoryEntry struct {
		Email        string
		Password     string
		CategoryName string
	}
	if err := ctx.ShouldBindBodyWith(&categoryEntry, binding.JSON); err != nil {
		ctx.JSON(400, helpers.CreateErrorResponse("BADi_REQUEST", err.Error()))
		return
	}
	user := models.User{Email: categoryEntry.Email, Password: categoryEntry.Password}
	category := models.Category{CategoryName: categoryEntry.CategoryName}
	code, msg := user.AuthenticateUser()
	if code == "AUTHENTICATED" {
		createCode, createMsg := category.CreateCategory(&user)
		if createCode == "SUCCESS" {
			ctx.JSON(200, helpers.CreateSuccessResponse(createCode, createMsg))
		} else if createCode == "CATEGORY_EXISTS" || createCode == "INVALID_CATEGORY" {
			ctx.JSON(401, helpers.CreateErrorResponse(createCode, createMsg))
		} else { //CreateCode is "DB_CONNECTIVITY_ISSUE" or "DB_INSERT_ERROR"  or anything
			ctx.JSON(500, helpers.CreateErrorResponse(createCode, createMsg))
		}
		return
	} else if code == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, helpers.CreateErrorResponse(code, msg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, helpers.CreateErrorResponse(code, msg))
	}
}

func DeleteCategory(ctx *gin.Context) {
	var categoryEntry struct {
		Email        string
		Password     string
		CategoryName string
	}
	if err := ctx.ShouldBindBodyWith(&categoryEntry, binding.JSON); err != nil {
		ctx.JSON(400, helpers.CreateErrorResponse("BAD_REQUEST", err.Error()))
		return
	}
	user := models.User{Email: categoryEntry.Email, Password: categoryEntry.Password}
	category := models.Category{CategoryName: categoryEntry.CategoryName}

	authCode, authMsg := user.AuthenticateUser()
	if authCode == "AUTHENTICATED" {
		if deleteCode, err := category.DeleteCategory(&user); err != nil {
			ctx.JSON(500, helpers.CreateErrorResponse(deleteCode, err.Error()))
		} else {
			ctx.JSON(200, helpers.CreateSuccessResponse(deleteCode, "Delete successful"))
		}

	} else if authCode == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, helpers.CreateErrorResponse(authCode, authMsg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, helpers.CreateErrorResponse(authCode, authMsg))
	}
}

func NewTransaction(ctx *gin.Context) {
	var transactionEntry struct {
		Email       string
		Password    string
		Time        time.Time
		Amount      int
		Category    string
		Description string
		SplitTag    string
		MapUrl      string
	}
	if err := ctx.ShouldBindBodyWith(&transactionEntry, binding.JSON); err != nil {
		ctx.JSON(400, helpers.CreateErrorResponse("BAD_REQUEST", err.Error()))
		return
	}

	fmt.Println(transactionEntry)
	user := models.User{Email: transactionEntry.Email, Password: transactionEntry.Password}
	authCode, authMsg := user.AuthenticateUser()
	if authCode == "AUTHENTICATED" {
		transaction := models.Transaction{UserId: user.ID, Time: transactionEntry.Time, Amount: transactionEntry.Amount, Category: transactionEntry.Category, Description: transactionEntry.Description, SplitTag: transactionEntry.SplitTag, MapUrl: transactionEntry.MapUrl}
		if msg, err := transaction.NewTransaction(); err != nil {
			ctx.JSON(500, helpers.CreateErrorResponse("INSERT", err.Error()))
		} else {
			ctx.JSON(200, helpers.CreateSuccessResponse("SUCCESS", msg))
		}
	} else if authCode == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, helpers.CreateErrorResponse(authCode, authMsg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, helpers.CreateErrorResponse(authCode, authMsg))
	}
}

func GetLastNTransactions(ctx *gin.Context) {
	var transactionFilter struct {
		Email    string
		Password string
		LastN    int
	}
	if err := ctx.ShouldBindBodyWith(&transactionFilter, binding.JSON); err != nil {
		ctx.JSON(400, helpers.CreateErrorResponse("BAD_REQUEST", err.Error()))
		return
	}
	user := models.User{Email: transactionFilter.Email, Password: transactionFilter.Password}
	authCode, authMsg := user.AuthenticateUser()
	if authCode == "AUTHENTICATED" {
		transactions, err := user.GetLastNTransactions(transactionFilter.LastN)
		fmt.Println(transactions)
		if err != nil {
			ctx.JSON(200, err.Error())
		} else {
			ctx.JSON(500, transactions)
		}
	} else if authCode == "INVALID_USERID_PASSWORD" {
		ctx.JSON(401, helpers.CreateErrorResponse(authCode, authMsg))
	} else { //DB_CONNECTIVITY_ISSUE or anything
		ctx.JSON(500, helpers.CreateErrorResponse(authCode, authMsg))
	}
}

func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}

func Home(ctx *gin.Context) {
	ctx.JSON(200, "Hello Karchu")
}
