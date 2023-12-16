package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId           uint
	CategoryId       uint
	Amount           int
	Time             time.Time
	Description      string
	SplitTag         string
	CategoryMappings []CategoryTransactionMapping `gorm:"foreignKey:TransactionId"`
}
