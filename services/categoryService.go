package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
)

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
