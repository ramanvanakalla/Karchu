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

func TransactionAlreadySplit(transactionId uint) (bool, error) {
	var transaction models.Transaction
	err := initializers.DB.Preload("Splits").First(&transaction, transactionId).Error
	if err != nil {
		return true, err
	}
	if len(transaction.Splits) > 0 {
		return true, nil
	}
	return false, nil
}

func SettleTransaction(userId uint, splitTransactionId uint) error {
	var splitTransaction models.SplitTransaction
	err := initializers.DB.First(&splitTransaction, splitTransactionId).Error
	if err != nil {
		return err
	}
	var transaction models.Transaction
	err = initializers.DB.Preload("CategoryMappings").First(&transaction, splitTransaction.SourceTransactionId).Error
	if err != nil {
		return err
	}
	categoryId := transaction.CategoryMappings[0].CategoryId
	SettledTransactionId, err := CreateTransactionV2(userId, time.Now(), -1*splitTransaction.Amount, categoryId, "settle trans", "SETTLED TRANS")
	if err != nil {
		return err
	}
	splitTransaction.SettledTransactionId = SettledTransactionId
	err = initializers.DB.Save(&splitTransaction).Error
	if err != nil {
		return err
	}
	return nil
}
