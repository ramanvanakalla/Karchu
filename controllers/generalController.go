package controllers

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.JSON(200, "Hello Karchu")
}
