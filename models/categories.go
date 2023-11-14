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
	category.lowerCategoryName()
	category.UserId = user.ID
	if isValid, validErr := category.validCategory(); isValid {
		if queryErr := initializers.DB.Where("user_id = ? and category_name = ?", user.ID, category.CategoryName).First(&category).Error; queryErr != nil {
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

func (category *Category) DeleteCategory(user *User) (string, error) {
	category.lowerCategoryName()
	category.UserId = user.ID
	if queryErr := initializers.DB.Where("user_id = ? and category_name = ?", user.ID, category.CategoryName).First(&category).Error; queryErr != nil {
		if queryErr == gorm.ErrRecordNotFound {
			return "CATEGORY_DOESNT_EXIST", queryErr
		} else {
			return "DB_ERROR", queryErr
		}
	} else {
		if deleteErr := initializers.DB.Delete(&category).Error; deleteErr != nil {
			return "DB_DELETE_ERROR", deleteErr
		} else {
			return "SUCCESS", nil
		}
	}

}
