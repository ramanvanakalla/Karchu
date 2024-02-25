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

// CreateUser godoc
// @Summary      create a user
// @Description  creates a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateUserReq true "enter Email, Name and Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /v1/user/ [post]
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
// @Router       /v1/user/auth [post]
func AuthUser(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	ctx.JSON(http.StatusOK, userIDUint)
}

// GetTransactionsListAsStringV2 godoc
// @Summary      Get transactions list as string of user
// @Description  returns transactions as string for UI
// @Tags         Transactions, V2
// @Accept       json
// @Produce      json
// @Param        request body requests.GetTransactionsReq true "enter Email,Password"
// @Success      200  {array} string "returns transaction strings as list"
// @Router       /v1/transactions/all [post]
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
	transactionList, ex := services.GetTransactionsListV2(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}

// GetTransactionsV2 godoc
// @Summary      Get transactions of user
// @Description  returns transactions
// @Tags         Transactions, V2
// @Accept       json
// @Produce      json
// @Param        request body requests.GetTransactionsReq true "enter Email,Password"
// @Success      200
// @Router       /v1/transactions/get [post]
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
	transactionList, ex := services.GetTransactionsV2(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}

// GetTransactionsV2 godoc
// @Summary      Get transactions of user
// @Description  returns transactions
// @Tags         Transactions, V2
// @Accept       json
// @Produce      json
// @Param        request body requests.GetFilteredTransactionsReq true "enter Email,Password"
// @Success      200
// @Router       /v1/transactions/GetTransactionsFiltered [post]
func GetTransactionsFiltered(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.GetFilteredTransactionsReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, req.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_START_DATE", err.Error()))
		ctx.Abort()
		return
	}
	EndDate, err := time.Parse(layout, req.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_END_DATE", err.Error()))
		ctx.Abort()
		return
	}
	transactionList, ex := services.GetTransactionsFiltered(userIDUint, startDate, EndDate, req.Categories, req.SplitTag)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionList)
}
