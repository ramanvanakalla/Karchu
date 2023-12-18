package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserId              uint                         `gorm:"uniqueIndex:idx_member"`
	CategoryName        string                       `gorm:"uniqueIndex:idx_member"`
	Transactions        []Transaction                `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE;"`
	TransactionMappings []CategoryTransactionMapping `gorm:"foreignKey:CategoryId"`
}
