package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/views"
	"fmt"
	"strings"
)

func validateAndNormalizeCategory(categoryName *string) bool {
	*categoryName = strings.TrimSpace(*categoryName)
	if len(*categoryName) < 1 {
		return false
	} else {
		return true
	}
}

func GetCategoriesByUserID(userID uint) ([]string, *exceptions.GeneralException) {
	categories, err := dao.GetCategoriesByUserID(userID)
	categoryArr := make([]string, 0)
	if err != nil {
		return categoryArr, exceptions.InternalServerError(err.Error(), "DB_GET_CATEGORY_ERROR")
	}
	for _, categoty := range categories {
		categoryArr = append(categoryArr, categoty.CategoryName)
	}
	return categoryArr, nil
}

func CreateCategoryForUserID(userId uint, categoryName string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&categoryName) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", categoryName), "INVALID_CATEGORY_FORMAT")
	}
	categoryId, err := dao.CreateCategory(userId, categoryName)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return categoryId, nil
}

func DeleteCategoryForUserID(userId uint, categoryName string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&categoryName) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", categoryName), "INVALID_CATEGORY_FORMAT")
	}
	delCategoryId, err := dao.DeleteCategory(userId, categoryName)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_DELETION_FAIL")
	}
	return delCategoryId, nil
}

func GetTransactionsOfCategory(userId uint, categoryName string) ([]string, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&categoryName) {
		return nil, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", categoryName), "INVALID_CATEGORY_FORMAT")
	}
	transactionsOfCategory, err := dao.GetTransactionsOfCategory(userId, categoryName)
	if err != nil {
		return nil, exceptions.InternalServerError(err.Error(), "TRANSACTIONS_GET_FAIL")
	}
	transactionsList := make([]string, 0)
	for _, transaction := range transactionsOfCategory {
		transactionView := views.NewTransactionWithCategory(transaction, categoryName)
		transactionsList = append(transactionsList, transactionView.ToString())
	}
	return transactionsList, nil
}
