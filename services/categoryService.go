package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/views"
	"fmt"
	"sort"
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

func GetTransactionsOfCategoryV2(userId uint, categoryName string) ([]string, *exceptions.GeneralException) {
	categoryName = strings.Split(categoryName, "-")[0]
	if !validateAndNormalizeCategory(&categoryName) {
		return nil, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", categoryName), "INVALID_CATEGORY_FORMAT")
	}
	transactions, err := dao.GetTransactionsOfCategoryV2(userId, categoryName)
	if err != nil {
		return nil, exceptions.InternalServerError(err.Error(), "TRANSACTIONS_GET_FAIL")
	}
	transactionsList := make([]string, 0)
	for _, transactionView := range transactions {
		transactionsList = append(transactionsList, transactionView.ToString())
	}
	return transactionsList, nil
}

func GetTransactionsOfCategory(userId uint, categoryName string) ([]string, *exceptions.GeneralException) {
	categoryName = strings.Split(categoryName, "-")[0]
	if !validateAndNormalizeCategory(&categoryName) {
		return nil, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", categoryName), "INVALID_CATEGORY_FORMAT")
	}
	transactionsOfCategory, err := dao.GetTransactionsOfCategory(userId, categoryName)
	if err != nil {
		return nil, exceptions.InternalServerError(err.Error(), "TRANSACTIONS_GET_FAIL")
	}
	transactionsList := make([]string, 0)
	transactionViewList := make([]views.TransactionWithCategory, 0)
	for _, transaction := range transactionsOfCategory {
		transactionViewList = append(transactionViewList, views.NewTransactionWithCategory(transaction, categoryName))
	}

	sort.Slice(transactionViewList, func(i, j int) bool {
		return transactionViewList[i].ID > transactionViewList[j].ID
	})
	for _, transactionView := range transactionViewList {
		transactionsList = append(transactionsList, transactionView.ToString())
	}
	return transactionsList, nil
}

func RenameCategory(userId uint, oldCategoryName string, newCategoryName string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&newCategoryName) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", newCategoryName), "INVALID_CATEGORY_FORMAT")
	}
	id, err := dao.RenameCategory(userId, oldCategoryName, newCategoryName)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "RENAME_FAIL")
	}
	return id, nil
}

func MergeCategory(userId uint, sourceCategoryName string, destinationCategoryName string) *exceptions.GeneralException {
	sourceCategoryId, sourceCaterr := dao.GetCategoryIdByUserIdAndCategoryName(userId, sourceCategoryName)
	if sourceCaterr != nil {
		return exceptions.InternalServerError(fmt.Sprintf("Failed to get CategoryId for %s", sourceCategoryName), "FAIL_CATEGORY_ID")
	}
	destCategoryId, destCaterr := dao.GetCategoryIdByUserIdAndCategoryName(userId, destinationCategoryName)
	if destCaterr != nil {
		return exceptions.InternalServerError(fmt.Sprintf("Failed to get CategoryId for %s", destinationCategoryName), "FAIL_CATEGORY_ID")
	}
	mergeErr := dao.MergeCategory(userId, sourceCategoryId, destCategoryId)
	if mergeErr != nil {
		return exceptions.InternalServerError(fmt.Sprintf("Failed to merge category %s into %s", sourceCategoryName, destinationCategoryName), "FAIL_CATEGORY_MERGE")
	}
	return nil
}
