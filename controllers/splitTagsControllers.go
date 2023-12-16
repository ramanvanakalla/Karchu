package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetSplitTags godoc
// @Summary      Get split tags
// @Description  returns split tags
// @Tags         Split-tags
// @Accept		 json
// @Produce      json
// @Success      200  {array} string "returns split tags"
// @Router       /split-tags [get]
func GetSplitTags(ctx *gin.Context) {
	ctx.JSON(200, []string{"No", "will split", "done splitting"})
}
