package views

import (
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

func (transaction *TransactionWithCategory) ToString() string {
	return fmt.Sprintf("Id: %d|Amount: %d|Category: %s|splitTag: %s|Desc: %s", transaction.ID, transaction.Amount, transaction.CategoryName, transaction.SplitTag, transaction.Description)
}
