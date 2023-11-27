package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

func AuthenticateUser(email string, password string) (uint, *exceptions.GeneralException) {
	userID, err := dao.GetUserId(email, password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorMessage := fmt.Sprintf("%s and %s doesn't exists", email, password)
		return userID, exceptions.UnAuthorizedError(errorMessage, "EMAIL_PWD_DOESNT_EXISTS")
	} else if !errors.Is(err, nil) {
		return userID, exceptions.InternalServerError(err.Error(), "DB_AUTH_ERROR")
	} else {
		return userID, nil
	}
}

func validateAndNormalizeEmail(email *string) bool {
	*email = strings.TrimSpace(*email)
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRe := regexp.MustCompile(emailRegex)
	return emailRe.MatchString(*email)
}

func validateAndNormalizePassword(password *string) bool {
	*password = strings.TrimSpace(*password)
	if len(*password) < 8 || !regexp.MustCompile(`[A-Z]`).MatchString(*password) || !regexp.MustCompile(`[a-z]`).MatchString(*password) || !regexp.MustCompile(`\d`).MatchString(*password) || !regexp.MustCompile(`[@#$%^&+=!]`).MatchString(*password) {
		return false
	} else {
		return true
	}
}

func validateAndNormalizeName(name *string) bool {
	*name = strings.TrimSpace(*name)
	nameRegex := `^[A-Za-z0-9_ ]{3,20}$`
	nameRe := regexp.MustCompile(nameRegex)
	return nameRe.MatchString(*name)
}

func CreateUser(email string, password string, name string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeEmail(&email) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid email %s format", email), "INVALID_EMAIL_FORMAT")
	}
	if !validateAndNormalizePassword(&password) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid password %s format", password), "INVALID_PASSWORD_FORMAT")
	}
	if !validateAndNormalizeName(&name) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid name %s format", name), "INVALID_NAME_FORMAT")
	}
	userId, err := dao.CreateUser(email, password, name)
	if err != nil {
		return userId, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return userId, nil
}
