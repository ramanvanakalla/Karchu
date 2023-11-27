package controllers

import (
	"Karchu/helpers"
	"Karchu/requests"
	"Karchu/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func CreateCategory(ctx *gin.Context) {
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
	var req requests.CreateCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.CreateCategoryForUserID(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, helpers.CreateSuccessResponse("CATEGORY_CREATED", fmt.Sprintf("category Id %d created", categoryId)))
}

func DeleteCategoryV2(ctx *gin.Context) {
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
	var req requests.DeleteCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("AUTH_BAD_REQUEST", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.DeleteCategoryForUserID(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, helpers.CreateSuccessResponse("CATEGORY_DELETED", fmt.Sprintf("category Id %d deleted", categoryId)))
}
