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
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
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
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
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

func DeleteCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
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

func GetTransactionOfCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.TransactionsOfCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionsOfCategory, ex := services.GetTransactionsOfCategory(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionsOfCategory)
}

func RenameCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.RenameCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.RenameCategory(userIDUint, req.OldCategoryName, req.NewCategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, helpers.CreateSuccessResponse("CATEGORY_RENAMED", fmt.Sprintf("Category ID %d renamed", categoryId)))
}

func MergeCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, helpers.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.MergeCategory
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.MergeCategory(userIDUint, req.SourceCategoryName, req.DestinationCategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, helpers.CreateErrorResponse(ex.Status, ex.Message))
	}
	ctx.JSON(http.StatusOK, helpers.CreateSuccessResponse("CATEGORY_MERGED", fmt.Sprintf("Category %s merged into %s", req.SourceCategoryName, req.DestinationCategoryName)))
}
