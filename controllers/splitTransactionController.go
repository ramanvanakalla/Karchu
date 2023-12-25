package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SplitTransactionWithOneFriend godoc
// @Summary      split a transaction
// @Description  split a transaction
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.SplitWithOneFriendReq true "split with one"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction/one [post]
func SplitTransactionWithOneFriend(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.SplitWithOneFriendReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.SplitTransactionWithOneFriend(userIDUint, req.TransString, req.Friend, req.Amount)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SPLIT", "Transaction got split"))
}

// GetUnSettledSplitTransactions godoc
// @Summary      Get splits of user
// @Description  Get splits of user
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.GetUnSettledSplitTransactionsReq true "get split transaction"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction/unsettled-splits [post]
func GetUnSettledSplitTransactions(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetUnSettledSplitTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	splitStrings, ex := services.GetUnSettledSplitTransactions(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, splitStrings)
}

// GetSettledSplitTransactions godoc
// @Summary      Get splits of user which are settled
// @Description  Get splits of user which are settled
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.GetSettledSplitTransactionsReq true "get split transaction"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction/settled-splits [post]
func GetSettledSplitTransactions(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetSettledSplitTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	splitStrings, ex := services.GetSettledSplitTransactions(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, splitStrings)
}

// GetSplitTransactions godoc
// @Summary      Get splits of user
// @Description  Get splits of user
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.GetSplitTransactionsReq true "get split transaction"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction/splits [post]
func GetSplitTransactions(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetSplitTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	splitStrings, ex := services.GetSplitTransactions(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, splitStrings)
}

// SplitTransaction godoc
// @Summary      split a transaction
// @Description  split a transaction
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.SplitTransactionReq true "split transaction"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction [post]
func SplitTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.SplitTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.SplitTransaction(userIDUint, uint(req.TransactionId), req.Splits)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SPLIT", "Transaction got split"))
}

// DeleteSplitTransaction godoc
// @Summary      deletes a alreadt split transaction
// @Description  deletes a alreadt split transaction
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.DeleteSplitTransactionReq true "delete split"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction [delete]
func DeleteSplitTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteSplitTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.DeleteSplitTransaction(userIDUint, req.TransactionId)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SPLIT_DELETE", "Transaction splits go deleted"))
}

// DeleteSplitTransactionString godoc
// @Summary      deletes a alreadt split transaction from transString
// @Description  deletes a alreadt split transaction from transString
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.DeleteSplitTransactionStringReq true "delete split"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction/str [delete]
func DeleteSplitTransactionString(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteSplitTransactionStringReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.DeleteSplitTransactionString(userIDUint, req.TransString)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SPLIT_DELETE", "Transaction splits go deleted"))
}

// settleSplitTransaction godoc
// @Summary      settle a split
// @Description  settle a split transaction
// @Tags         settleTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.SettleTransactionReq true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /settle-transaction [post]
func SettleSplitTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.SettleTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.SettleTransaction(userIDUint, req.SplitTransactionId)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SETTLED", "Transaction got settled"))
}

// SettleSplitTransactionString godoc
// @Summary      settle a split
// @Description  settle a split transaction
// @Tags         settleTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.SettleTransactionReq true "settle transaction"
// @Success      200  {object} responses.SuccessRes
// @Router       /settle-transaction/str [post]
func SettleSplitTransactionString(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.SettleTransactionStringReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.SettleTransactionString(userIDUint, req.SplitTransactionString)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_SETTLED", "Transaction got settled"))
}

// UnsettleSplitTransaction godoc
// @Summary      settle a split
// @Description  settle a split transaction
// @Tags         settleTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.UnSettleTransactionReq true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /settle-transaction [delete]
func UnSettleSplitTransaction(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.UnSettleTransactionReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.UnSettleTransaction(userIDUint, req.SplitTransactionId)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_UNSETTLED", "Transaction got un-settled"))
}

// UnSettleSplitTransactionString godoc
// @Summary      settle a split
// @Description  settle a split transaction
// @Tags         settleTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.UnSettleTransactionReq true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /settle-transaction [delete]
func UnSettleSplitTransactionString(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.UnSettleTransactionStringReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.UnSettleTransactionString(userIDUint, req.SplitTransactionString)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("TRANSACTION_UNSETTLED", "Transaction got un-settled"))
}

// MoneyLentFriend godoc
// @Summary      Money Lent to a friend
// @Description  Money lent to a friend
// @Tags         MoneyLent
// @Accept       json
// @Produce      json
// @Param        request body requests.MoneyLentFriend true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /money-lent [POST]
func MoneyLentToFriend(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.MoneyLentFriend
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	moneyLentByCategory, ex := services.MoneyLentToFriend(userIDUint, req.FriendName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, moneyLentByCategory)
}
