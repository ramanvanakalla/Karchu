package models

import (
	"Karchu/initializers"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Password     string
	Email        string `gorm:"unique"`
	Transactions []Transaction
	Categories   []Category
}

func (user *User) lowerUserName() {
	user.Name = strings.ToLower(user.Name)
}

func (user *User) isValidPassword() bool {
	if len(user.Password) < 8 || !regexp.MustCompile(`[A-Z]`).MatchString(user.Password) || !regexp.MustCompile(`[a-z]`).MatchString(user.Password) || !regexp.MustCompile(`\d`).MatchString(user.Password) || !regexp.MustCompile(`[@#$%^&+=!]`).MatchString(user.Password) {
		return false
	} else {
		return true
	}
}

func (user *User) isValidEmail() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRe := regexp.MustCompile(emailRegex)
	return emailRe.MatchString(user.Email)
}

func (user *User) isValidName() bool {
	nameRegex := `^[a-z0-9_ ]{3,20}$`
	nameRe := regexp.MustCompile(nameRegex)
	return nameRe.MatchString(user.Name)
}

func (user *User) validate() error {
	var err error
	if !user.isValidName() {
		err = fmt.Errorf("%s is not a valid name", user.Name)
	} else if !user.isValidEmail() {
		err = fmt.Errorf("%s is not a valid email address", user.Email)
	} else if !user.isValidPassword() {
		err = fmt.Errorf("%s is not a valid password", user.Password)
	} else {
		err = nil
	}
	return err
}

func (user *User) alreadyExists() (bool, error) {
	var queryUser User
	err := initializers.DB.First(&queryUser, "name = ?", user.Name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		} else {
			return true, err
		}
	} else {
		return false, nil
	}

}

func (user *User) CreateUser() (string, string) {
	user.lowerUserName()
	if err := user.validate(); err != nil {
		return "INVALID_USER_DETAILS", err.Error()
	}
	if alreadyExists, err := user.alreadyExists(); err != nil {
		return "DB_CONNECTIVITY_ISSUE", err.Error()
	} else {
		if !alreadyExists {
			return "EMAIL_ALREADY_EXISTS", fmt.Sprintf("%s is already in use", user.Email)
		} else {
			if err := initializers.DB.Create(&user).Error; err != nil {
				return "DB_INSERT_ERROR", err.Error()
			} else {
				return "SUCCESS", fmt.Sprintf("user Id %d is created", user.ID)
			}
		}
	}

}