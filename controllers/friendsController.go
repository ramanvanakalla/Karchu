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

// CreateFriend godoc
// @Summary      create a friend
// @Description  creates a friend
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateFriendReq true "enter Email, Password and friend Name"
// @Success      200  {object} responses.SuccessRes
// @Router       /v2/friends/ [post]
func CreateFriend(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateFriendReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	friendId, ex := services.CreateFriend(userIDUint, req.FriendName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("FRIEND_CREATED", fmt.Sprintf("Friend %s with id %d is created", req.FriendName, friendId)))
}

// GetFriend godoc
// @Summary      get friends
// @Description  get friends
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        request body requests.GetFriendsReq true "enter Email, Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /v2/friends/get [post]
func GetFriends(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateFriendReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	friends, ex := services.GetFriends(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, friends)
}

// GetFriend godoc
// @Summary      get money friends
// @Description  get money friends
// @Tags         Friends
// @Accept       json
// @Produce      json
// @Param        request body requests.MoneyFriends true "enter Email, Password"
// @Success      200  {object} responses.SuccessRes
// @Router       /v2/friends/money-friends [post]
func moneyFriends(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateFriendReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	friends, ex := services.MoneyInvolvedWithFriends(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, friends)
}
