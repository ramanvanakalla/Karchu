package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/views"
	"time"
)

func CreateTransaction(userId uint, time time.Time, amount int, category string, categoryId uint, description string, splitTag string) (uint, error) {
	transaction := models.Transaction{UserId: userId, Time: time, Amount: amount, CategoryId: categoryId, Description: description, SplitTag: splitTag}
	err := initializers.DB.Create(&transaction).Error
	return transaction.ID, err
}

func DeleteTransactionbyTransactionIdAndUserId(transactionId uint, userId uint) (uint, error) {
	var transaction models.Transaction
	err := initializers.DB.
		Model(&models.Transaction{}).
		Where("id = ? and user_id = ?", transactionId, userId).
		First(&transaction).
		Error
	if err != nil {
		return 0, err
	}
	deletionErr := initializers.DB.Delete(&transaction).Error
	return transaction.ID, deletionErr
}

func GetLastNTransactionsByUserId(userId uint, lastN int) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	err := initializers.DB.
		Model(&models.Transaction{}).
		Select("transactions.*,categories.category_name").
		Joins("JOIN categories ON transactions.category_id = categories.id").
		Where("transactions.user_id = ?", userId).
		Order("transactions.created_at desc").
		Limit(lastN).
		Find(&transactions).
		Error
	return transactions, err
}

func GetNetMoneySpentByCategory(userID uint) ([]views.NetCategorySum, error) {
	var amountByCategory []views.NetCategorySum
	err := initializers.DB.
		Model(&models.Transaction{}).
		Select("category, sum(amount) as net_amount").
		Where("user_id = ?", userID).
		Group("category").
		Order("net_amount desc").
		Scan(&amountByCategory).
		Error
	return amountByCategory, err
}
