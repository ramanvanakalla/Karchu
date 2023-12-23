package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateTransactionV2 godoc
// @Summary      creates a transaction for a user V2
// @Description  create a transaction with category V2
// @Tags         Transactions, V2
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateTransactionReq true "enter Email,Password"
// @Success      200  {array} responses.SuccessRes
// @Router       /transactions [post]
func NewTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionId, ex := services.CreateTransactionV2(userIDUint, req.Amount, req.Category, req.Description, req.SplitTag)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_CREATED", fmt.Sprintf("transaction Id %d created", transactionId)))
}

// CreateTransactionAndSplitWithOne godoc
// @Summary      creates a transaction and also split
// @Description  create a transaction and also split
// @Tags         trans-split-with-one
// @Accept       json
// @Produce      json
// @Param        request body requests.TransactionAndSplitWithOneReq true "add transaction and also split"
// @Success      200  {array} responses.SuccessRes
// @Router       /v2/trans-split-with-one [post]
func CreateTransactionAndSplitWithOne(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.TransactionAndSplitWithOneReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	if ex := services.CreateTransactionAndSplitWithOne(userIDUint, time.Now(), req.Amount, req.Category, req.Description, req.SplitTag, req.FriendName, req.SplitAmount); ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_AND_SPLT_CREATED", "Transaction and split succesfully created"))
}

// DeleteTransactionFromTransString godoc
// @Summary      delete transaction for a user
// @Description  delete a transaction for a given trans string
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.DeleteTransactionFromTransStringReq true "enter Email,Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /transactions/str [delete]
func DeleteTransactionFromTransString(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteTransactionFromTransStringReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	delTransactionId, ex := services.DeleteTransactionFromTransString(req.TransString, userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANS_DELETED", fmt.Sprintf("Trans Id %d deleted", delTransactionId)))
}

// DeleteTransaction godoc
// @Summary      delete transaction for a user
// @Description  delete a transaction
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.DeleteTransactionReq true "enter Email,Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /transactions [delete]
func DeleteTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	delTransactionId, ex := services.DeleteTransaction(req.TransactionId, userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANS_DELETED", fmt.Sprintf("Trans Id %d deleted", delTransactionId)))
}

// GetLastNTransactionV2 godoc
// @Summary      Get last N transactions of user
// @Description  Get last N transaction list of user
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.GetLastNTransactionsReq true "enter Email,Password"
// @Success      200  {array} string "last N transactions list"
// @Router       /transactions/last-n [post]
func GetLastNTransactions(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetLastNTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionList, ex := services.GetLastNTransactionsListV2(userIDUint, req.LastN)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}

// GetNetMoneySpentByCategory2 godoc
// @Summary      get money spent on each category
// @Description  get money spent on each category
// @Tags         Net-Amount
// @Accept       json
// @Produce      json
// @Param        request body requests.NetAmountByCategoryReq true "enter Email,Password"
// @Success      200  {array} string "money spent on each category as list"
// @Router       /net-amount/categories [post]
func GetNetMoneySpentByCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.NetAmountByCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	netByCategoriesList, ex := services.GetNetMoneySpentByCategory2(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, netByCategoriesList)
}
