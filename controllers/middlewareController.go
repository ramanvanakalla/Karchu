package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func skipAuth(ctx *gin.Context) bool {
	if ctx.Request.URL.Path == "/v1/user" && ctx.Request.Method == http.MethodPost {
		return true
	} else if ctx.Request.URL.Path == "/v1/split-tags" && ctx.Request.Method == http.MethodGet {
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
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	userID, ex := services.AuthenticateUser(userReq.Email, userReq.Password)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		ctx.Abort()
		return
	}
	ctx.Set("USER_ID", userID)
	ctx.Next()
}
