package controllers

import (
	"Karchu/helpers"
	"Karchu/requests"
	"Karchu/services"
	"fmt"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	fmt.Println("Received request: ", string(body))
	if skipAuth(ctx) {
		ctx.Next()
		return
	}
	var userReq requests.UserReq
	if err := ctx.ShouldBindBodyWith(&userReq, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
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
