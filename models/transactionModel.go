package models

import (
	"Karchu/initializers"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId      uint
	Amount      int
	Time        time.Time
	Category    string
	Description string
	SplitTag    string
	MapUrl      string
}

func (transaction *Transaction) NewTransaction() (string, error) {
	if err := initializers.DB.Create(&transaction).Error; err != nil {
		return "DB_INSERT_ERROR", err
	} else {
		return fmt.Sprintf("transaction id %d is created", transaction.ID), nil
	}
}
