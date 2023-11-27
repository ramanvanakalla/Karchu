package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
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
	fmt.Println(userId)
	fmt.Println(categoryName)
	delCategoryId, err := dao.DeleteCategory(userId, categoryName)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_DELETION_FAIL")
	}
	return delCategoryId, nil
}
