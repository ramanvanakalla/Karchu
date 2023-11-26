package controllers

import (
	"Karchu/helpers"
	"Karchu/requests"
	"Karchu/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func skipAuth(ctx *gin.Context) bool {
	if ctx.Request.URL.Path == "/v1/user" && ctx.Request.Method == http.MethodPost {
		return true
	} else {
		return false
	}
}

func AuthMiddleware(ctx *gin.Context) {
	if skipAuth(ctx) {
		ctx.Next()
		return
	}
	var userReq requests.UserReq
	if err := ctx.ShouldBindBodyWith(&userReq, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	userID, ex := services.AuthenticateUser(userReq.Email, userReq.Password)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		ctx.Abort()
		return
	}
	ctx.Set("USER_ID", userID)
	ctx.Next()
}
