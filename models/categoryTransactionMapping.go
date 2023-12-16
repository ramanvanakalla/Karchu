package models

import "gorm.io/gorm"

type CategoryTransactionMapping struct {
	gorm.Model
	TransactionId uint `gorm:"index;foreignkey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryId    uint `gorm:"index;foreignkey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
