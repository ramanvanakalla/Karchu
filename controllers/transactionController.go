package controllers

import (
	"Karchu/helpers"
	"Karchu/requests"
	"Karchu/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func NewTransactionV2(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	transactionId, ex := services.CreateTransaction(userIDUint, req.Time, req.Amount, req.Category, req.Description, req.SplitTag, req.MapUrl)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, helpers.CreateSuccessResponse("TRANSACTION_CREATED", fmt.Sprintf("transaction Id %d created", transactionId)))
}

func GetLastNTransactionsV2(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetLastNTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	transactionList, ex := services.GetLastNTransactionsList(userIDUint, req.LastN)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}
