package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/views"
	"errors"
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

func GetTransactionsOfCategoryV2(userId uint, categoryName string) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	err := initializers.DB.
		Table("categories").
		Where("categories.user_id = ? AND categories.category_name = ?", userId, categoryName).
		Joins("JOIN category_transaction_mappings ON categories.id = category_transaction_mappings.category_id").
		Joins("JOIN transactions ON category_transaction_mappings.transaction_id = transactions.id AND transactions.deleted_at is NULL").
		Select("transactions.*, categories.category_name").
		Order("transaction_id desc").
		Find(&transactions).
		Error
	return transactions, err
}

func RenameCategory(userId uint, oldCategoryName string, newCategoryName string) (uint, error) {
	err := initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ? AND category_name = ?", userId, oldCategoryName).
		Update("category_name", newCategoryName).
		Error
	if err != nil {
		return 0, err
	}
	var updatedCategory models.Category
	err = initializers.DB.
		Model(&models.Category{}).
		Where("user_id = ? AND category_name = ?", userId, newCategoryName).
		First(&updatedCategory).
		Error
	return updatedCategory.ID, err
}

func MergeCategory(userId uint, sourceCategoryId uint, destinationCategoryId uint) error {
	if sourceCategoryId == destinationCategoryId {
		return errors.New("source and destination category are same")
	}
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Update Transactions
	updateErr := tx.Model(&models.Transaction{}).
		Where("category_id = ?", sourceCategoryId).
		Update("category_id", destinationCategoryId).
		Error
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	// Step 2: Delete Source Category
	deleteErr := tx.Delete(&models.Category{}, sourceCategoryId).Error
	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	// Commit the transaction if everything is successful
	return tx.Commit().Error
}

func MergeCategoryV2(userId uint, sourceCategoryId uint, destinationCategoryId uint) error {
	if sourceCategoryId == destinationCategoryId {
		return errors.New("source and destination category are same")
	}
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Update Transactions
	updateErr := tx.Model(&models.CategoryTransactionMapping{}).
		Where("category_id = ?", sourceCategoryId).
		Update("category_id", destinationCategoryId).
		Error
	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	// Step 2: Delete Source Category
	deleteErr := tx.Delete(&models.Category{}, sourceCategoryId).Error
	if deleteErr != nil {
		tx.Rollback()
		return deleteErr
	}

	// Commit the transaction if everything is successful
	return tx.Commit().Error
}
