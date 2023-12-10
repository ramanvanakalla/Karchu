package dao

import (
	"Karchu/initializers"
	"Karchu/models"
)

func GetUserId(email string, password string) (uint, error) {
	var user models.User
	err := initializers.DB.
		Model(&models.User{}).
		Where("email = ? AND password = ?", email, password).
		First(&user).
		Error

	return user.ID, err
}

func CreateUser(email string, password string, name string) (uint, error) {
	newUser := models.User{Name: name, Password: password, Email: email}
	err := initializers.DB.Create(&newUser).Error
	return newUser.ID, err
}

func GetAllTransactionsByUserId(userId uint) (map[string][]models.Transaction, error) {
	transactionsByCategory := make(map[string][]models.Transaction)

	var user models.User
	err := initializers.DB.
		Preload("Categories").
		Preload("Categories.Transactions").
		First(&user, userId).
		Error
	if err != nil {
		return nil, err
	}
	for _, category := range user.Categories {
		transactionsByCategory[category.CategoryName] = category.Transactions
	}
	return transactionsByCategory, nil
}
