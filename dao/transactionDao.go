package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"time"
)

func CreateTransaction(userId uint, time time.Time, amount int, category string, categoryId uint, description string, splitTag string, mapUrl string) (uint, error) {
	transaction := models.Transaction{UserId: userId, Time: time, Amount: amount, Category: category, CategoryId: categoryId, Description: description, SplitTag: splitTag, MapUrl: mapUrl}
	err := initializers.DB.Create(&transaction).Error
	return transaction.ID, err
}

func GetLastNTransactionsByUserId(userId uint, lastN int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := initializers.DB.
		Model(models.Transaction{}).
		Where("user_id = ?", userId).
		Limit(lastN).
		Find(&transactions).
		Error
	return transactions, err
}
