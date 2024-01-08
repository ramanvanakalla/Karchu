package controllers

import (
	"Karchu/requests"
	"Karchu/responses"
	"Karchu/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func addExtraCategoriesForUI(categoriesArr *[]string) {
	*categoriesArr = append(*categoriesArr, "Add a New Category?")
}

// GetCategories godoc
// @Summary      Get categories of user
// @Description  returns array of categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        request body requests.UserReq true "enter Email,Password"
// @Success      200  {array} string "list of categories"
// @Router       /categories/all [post]
// @Router		 /v1/categories/n [post]
func GetCategories(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	categoriesArr, ex := services.GetCategoriesByUserID(userIDUint)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Message, ex.Status))
		return
	}

	routePath := ctx.Request.URL.Path
	if routePath == "/v1/categories/n" {
		addExtraCategoriesForUI(&categoriesArr)
	}

	ctx.JSON(http.StatusOK, categoriesArr)
}

// CreateCategory godoc
// @Summary      create a category
// @Description  creates a category for a user
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateCategoryReq true "enter credentials"
// @Success      200  {object} responses.SuccessRes
// @Router       /v1/categories/ [post]
func CreateCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.CreateCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.CreateCategoryForUserID(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("CATEGORY_CREATED", fmt.Sprintf("category Id %d created", categoryId)))
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Deletes a category for a user
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        request body requests.DeleteCategoryReq true "enter credentials"
// @Success      200  {object} responses.SuccessRes
// @Router       /v1/categories/ [delete]
func DeleteCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.DeleteCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.DeleteCategoryForUserID(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("CATEGORY_DELETED", fmt.Sprintf("category Id %d deleted", categoryId)))
}

// GetTransactionOfCategoryV2 godoc
// @Summary      returns transactions of category
// @Description  returns transactions list for a category
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.TransactionsOfCategoryReq true "enter Email,Password"
// @Success      200  {array} string
// @Router       /v1/transactions/category [post]
func GetTransactionStringsOfCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.TransactionsOfCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionsOfCategory, ex := services.GetTransactionStringsOfCategoryV2(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionsOfCategory)
}

// GetTransactionOfCategory godoc
// @Summary      returns transactions of category
// @Description  returns transactions list for a category
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.TransactionsOfCategoryReq true "enter Email,Password"
// @Success      200 array views.TransactionWithCategory
// @Router       /v1/transactions/categories [post]
func GetTransactionsOfCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.TransactionsOfCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	transactionsOfCategory, ex := services.GetTransactionsOfCategoryV2(userIDUint, req.CategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, transactionsOfCategory)
}

// RenameCategory godoc
// @Summary      Rename a category
// @Description  Renames a category for a user
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        request body requests.RenameCategoryReq true "enter credentials"
// @Success      200  {object} responses.SuccessRes
// @Router       /v1/categories/ [patch]
func RenameCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.RenameCategoryReq
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	categoryId, ex := services.RenameCategory(userIDUint, req.OldCategoryName, req.NewCategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("CATEGORY_RENAMED", fmt.Sprintf("Category ID %d renamed", categoryId)))
}

// MergeCategoryV2 godoc
// @Summary      Merge a category into another
// @Description  Merges a category into another, all the transactions of soruce category will now be part of destination category
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        request body requests.MergeCategory true "enter credentials"
// @Success      200  {object} responses.SuccessRes
// @Router       /v1/categories/merge [post]
func MergeCategory(ctx *gin.Context) {
	userIDUint, ok := getUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, responses.CreateErrorResponse("Error while getting userId", "USERID_NOT_SET_CTX"))
		return
	}
	var req requests.MergeCategory
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.CreateErrorResponse("CANT_PARSE_REQ", err.Error()))
		ctx.Abort()
		return
	}
	ex := services.MergeCategoryV2(userIDUint, req.SourceCategoryName, req.DestinationCategoryName)
	if ex != nil {
		ctx.JSON(ex.StatusCode, responses.CreateErrorResponse(ex.Status, ex.Message))
		return
	}
	ctx.JSON(http.StatusOK, responses.CreateSuccessResponse("CATEGORY_MERGED", fmt.Sprintf("Category %s merged into %s", req.SourceCategoryName, req.DestinationCategoryName)))
}
