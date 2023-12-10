package dao

import (
	"Karchu/initializers"
	"Karchu/models"
)

func GetCategoriesByUserID(userID uint) ([]models.Category, error) {
	categories := make([]models.Category, 0)
	err := initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ?", userID).
		Find(&categories).
		Error
	return categories, err
}

func GetCategoryIdByUserIdAndCategoryName(userId uint, categoryName string) (uint, error) {
	var category models.Category
	err := initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ? and category_name = ?", userId, categoryName).
		First(&category).
		Error
	return category.ID, err
}

func GetCategoryNameFromId(categoryId uint) (string, error) {
	var category models.Category
	err := initializers.DB.
		Model(&models.Category{}).
		Where("id = ?", categoryId).
		First(&category).
		Error
	return category.CategoryName, err
}

func CreateCategory(userID uint, categoryName string) (uint, error) {
	newCategory := models.Category{UserId: userID, CategoryName: categoryName}
	err := initializers.DB.Create(&newCategory).Error
	return newCategory.ID, err
}

func DeleteCategory(userID uint, categoryName string) (uint, error) {
	var category models.Category
	err := initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ? AND category_name = ?", userID, categoryName).
		First(&category).
		Error
	if err != nil {
		return 0, err
	}
	deletionErr := initializers.DB.Delete(&category).Error
	return category.ID, deletionErr
}

func GetTransactionsOfCategory(userId uint, categoryName string) ([]models.Transaction, error) {
	var category models.Category
	err := initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ? AND category_name = ?", userId, categoryName).
		Preload("Transactions").
		First(&category).
		Error
	return category.Transactions, err
}
