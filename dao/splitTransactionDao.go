package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/requests"
	"fmt"
)

func AddSplitTransactions(userId uint, transactionId uint, splits []requests.FriendSplit) error {
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, split := range splits {
		splitTransaction := models.SplitTransaction{SourceTransactionId: transactionId, Amount: split.Amount, FriendId: uint(split.FriendId), SettledTransactionId: transactionId}
		fmt.Println(splitTransaction)
		err := initializers.DB.Create(&splitTransaction).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
