package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUser godoc
// @Summary      create a user
// @Description  creates a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateUserReq true "enter Email, Name and Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /user/ [post]
func CreateUser(ctx *gin.Context) {
	var req requests.CreateUserReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	userId, ex := services.CreateUser(req.Email, req.Password, req.Name)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("USER_CREATED", fmt.Sprintf("user id %d is created", userId)))
}

// AuthUser godoc
// @Summary      Authorizes the user creds
// @Description  returns userId
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body requests.UserReq true "enter Email and Password"
// @Success      200  {int} int "UserId"
// @Router       /user/auth [post]
func AuthUser(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	ctx.JSON(http.StatusOK, userIDUint)
}

// GetTransactionsListAsString godoc
// @Summary      Get transactions list as string of user
// @Description  returns transactions as string for UI
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.GetTransactionsReq true "enter Email,Password"
// @Success      200  {array} string "returns transaction strings as list"
// @Router       /transactions/all [post]
func GetTransactionsListOfUser(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionList, ex := services.GetTransactionsList(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}

// GetTransactions godoc
// @Summary      Get transactions of user
// @Description  returns transactions
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateTransactionReq true "enter Email,Password"
// @Success      200  {array} models.Transaction "returns transaction as list"
// @Router       /transactions/get [post]
func GetTransactions(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionList, ex := services.GetTransactions(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}
