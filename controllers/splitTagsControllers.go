package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}
