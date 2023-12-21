package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/requests"
	"errors"
	"fmt"
	"time"
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

func VerifySplitTransaction(userId uint, splitTransactionId uint) error {
	var splitTransaction models.SplitTransaction
	err := initializers.DB.First(&splitTransaction, splitTransactionId).Error
	if err != nil {
		return err
	}
	var transaction models.Transaction
	err = initializers.DB.First(&transaction, splitTransaction.SourceTransactionId).Error
	if err != nil {
		return err
	}
	if transaction.UserId != userId {
		return errors.New("split transaction doesn't exists for this user")
	}
	return nil
}

func SettleTransaction(userId uint, splitTransactionId uint) error {
	var splitTransaction models.SplitTransaction
	err := initializers.DB.First(&splitTransaction, splitTransactionId).Error
	if err != nil {
		return err
	}
	SettledTransactionId, err := CreateTransactionV2(userId, time.Now(), splitTransaction.Amount)
}
