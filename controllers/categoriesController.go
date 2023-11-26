package controllers

import (
	"Karchu/helpers"
	"Karchu/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addExtraCategoriesForUI(categoriesArr *[]string) {
	*categoriesArr = append(*categoriesArr, "New-Category")
}

func GetCategories(ctx *gin.Context) {
	userID, exists := ctx.Get("USER_ID")
	if !exists {
		errorMessage := "USER ID doesn't exist on gin context"
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(errorMessage, "USERID_NOT_SET_CTX"))
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		errorMessage := "UserId is not typecastable"
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse(errorMessage, "USERID_TYPECAST_FAILED"))
		return
	}
	categoriesArr, ex := services.GetCategoriesByUserID(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Message, ex.Status))
		return
	}

	routePath := ctx.Request.URL.Path
	if routePath == "/v1/categories/n" {
		addExtraCategoriesForUI(&categoriesArr)
	}

	ctx.JSON(http.StatusOK, categoriesArr)
}
