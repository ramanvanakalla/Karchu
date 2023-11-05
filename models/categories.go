package models

import (
	"Karchu/initializers"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserId       uint
	CategoryName string
}

func (category *Category) lowerCategoryName() {
	category.CategoryName = strings.ToLower(category.CategoryName)
}

func (category Category) validCategory() (bool, error) {
	if len(category.CategoryName) > 0 {
		return true, nil
	} else {
		return false, errors.New("Category is not valid")
	}
}

func (category Category) CreateCategory(user *User) (string, string) {
	category.UserId = user.ID
	category.lowerCategoryName()
	if isValid, validErr := category.validCategory(); isValid {
		if queryErr := initializers.DB.First(&category).Error; queryErr != nil {
			if queryErr == gorm.ErrRecordNotFound {
				return "CATEGORY_EXISTS", fmt.Sprintf("Category %s already exists", category.CategoryName)
			} else {
				return "DB_CONNECTIVITY_ISSUE", queryErr.Error()
			}
		} else {
			if err := initializers.DB.Create(category); err != nil {
				return "DB_INSERT_ERROR", validErr.Error()
			} else {
				return "SUCCESS", fmt.Sprintf("category Id %d is created", category.ID)
			}
		}
	} else {
		return "INVALID_USER_DETAILS", validErr.Error()
	}

}
