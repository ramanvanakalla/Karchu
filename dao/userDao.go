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

func GetAllTransactionsByUserId(userId uint) ([]models.Transaction, error) {
	var user models.User
	err := initializers.DB.Preload("Transactions").
		Preload("Categories").
		First(&user, userId).
		Error
	return user.Transactions, err
}
