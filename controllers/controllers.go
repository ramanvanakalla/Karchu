package controllers

import "github.com/gin-gonic/gin"

func GetCategories(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Food":         "",
		"Travel":       "",
		"healthy food": "",
		"Office food":  "",
	})
}

func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}
