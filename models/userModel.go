package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Password     string
	Email        string        `gorm:"unique"`
	Transactions []Transaction `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Categories   []Category    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Friends      []Friend      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}
