package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SplitTransaction godoc
// @Summary      split a transaction
// @Description  split a transaction
// @Tags         SplitTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.SplitTransactionReq true
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

// settleSplitTransaction godoc
// @Summary      settle a split
// @Description  settle a split transaction
// @Tags         settleTransaction
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateFriendReq true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /split-transaction [post]
func settleSplitTransaction(ctx *gin.Context) {
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
}
