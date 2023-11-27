package controllers

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.JSON(200, "Hello Karchu")
}

func getUserID(ctx *gin.Context) (uint, bool) {
	userID, exists := ctx.Get("USER_ID")
	if !exists {
		return 0, false
	}
	userIDUint, ok := userID.(uint)
	return userIDUint, ok
}
