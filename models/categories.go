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

func (category *Category) CreateCategory(user *User) (string, string) {
	category.UserId = user.ID
	category.lowerCategoryName()
	if isValid, validErr := category.validCategory(); isValid {
		if queryErr := initializers.DB.Where("user_id = ? and category_name = ?", category.UserId, category.CategoryName).First(&category).Error; queryErr != nil {
			if queryErr == gorm.ErrRecordNotFound {
				if createErr := initializers.DB.Create(&category).Error; createErr != nil {
					return "DB_INSERT_ERROR", createErr.Error()
				} else {
					return "SUCCESS", fmt.Sprintf("category Id %d is created", category.ID)
				}
			} else {
				return "DB_CONNECTIVITY_ISSUE", queryErr.Error()
			}
		} else {
			return "CATEGORY_EXISTS", fmt.Sprintf("Category %s already exists", category.CategoryName)
		}
	} else {
		return "INVALID_CATEGORY", validErr.Error()
	}
}
