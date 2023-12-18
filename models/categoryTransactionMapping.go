package models

import "gorm.io/gorm"

type CategoryTransactionMapping struct {
	gorm.Model
	TransactionId uint `gorm:"uniqueIndex:idx_transaction_category;index;foreignkey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryId    uint `gorm:"uniqueIndex:idx_transaction_category;index;foreignkey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
