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
