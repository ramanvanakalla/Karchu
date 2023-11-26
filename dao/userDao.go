package dao

import (
	"Karchu/initializers"
	"Karchu/models"
)

func GetUserId(email string, password string) (uint, error) {
	var userID uint
	err := initializers.DB.
		Model(&models.User{}).
		Where("email = ? AND password = ?", email, password).
		Pluck("id", &userID).
		Error

	return userID, err
}

func CreateUser(email string, password string, name string) (uint, error) {
	newUser := models.User{Name: name, Password: password, Email: email}
	err := initializers.DB.Create(&newUser).Error
	return newUser.ID, err
}
