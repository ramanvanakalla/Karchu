package models

import (
	"gorm.io/gorm"
)

type ModelSplitMapping struct {
	gorm.Model
	UserId         uint           `gorm:"uniqueIndex:idx_user_split_deleted;index"`
	ModelSplitName string         `gorm:"uniqueIndex:idx_user_split_deleted;index"`
	DeletedAt      gorm.DeletedAt `gorm:"uniqueIndex:idx_user_split_deleted;index"`
}
