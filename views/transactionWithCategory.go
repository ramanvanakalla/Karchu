package views

import (
	"Karchu/models"
	"fmt"
	"time"
)

type TransactionWithCategory struct {
	ID           uint
	CategoryName string
	Amount       int
	Time         time.Time
	Description  string
	SplitTag     string
}

func NewTransactionWithCategory(transaction models.Transaction, categoryName string) TransactionWithCategory {
	return TransactionWithCategory{ID: transaction.ID, CategoryName: categoryName, Amount: transaction.Amount, Time: transaction.Time, Description: transaction.Description, SplitTag: transaction.SplitTag}
}

func (transaction *TransactionWithCategory) ToString() string {
	return fmt.Sprintf("Id: %d|Amount: %d|Category: %s|splitTag: %s|Desc: %s", transaction.ID, transaction.Amount, transaction.CategoryName, transaction.SplitTag, transaction.Description)
}
