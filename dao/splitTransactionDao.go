package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/requests"
	"Karchu/views"
	"errors"
	"fmt"
	"sort"
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
		fmt.Println(split.FriendId)
		splitTransaction := models.SplitTransaction{SourceTransactionId: transactionId, Amount: split.Amount, FriendId: split.FriendId}
		fmt.Println(splitTransaction)
		err := tx.Create(&splitTransaction).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	var transaction models.Transaction
	if err := tx.First(&transaction, transactionId).Error; err != nil {
		tx.Rollback()
		return err
	}
	transaction.SplitTag = "Done Split"
	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return err
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
	if splitTransaction.SettledTransactionId != nil {
		return errors.New("transaction is already Split")
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
	splitTransaction.SettledTransactionId = &SettledTransactionId
	err = initializers.DB.Save(&splitTransaction).Error
	if err != nil {
		return err
	}
	return nil
}

func UnSettleSplitTransaction(userId uint, splitTransactionId uint) error {
	var splitTransaction models.SplitTransaction
	if err := initializers.DB.First(&splitTransaction, splitTransactionId).Error; err != nil {
		return err
	}
	var settledTransaction models.Transaction
	if err := initializers.DB.First(&settledTransaction, splitTransaction.SettledTransactionId).Error; err != nil {
		return err
	}
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Delete(&settledTransaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	splitTransaction.SettledTransactionId = nil
	if err := tx.Save(&splitTransaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteTransactionSplit(userId uint, transactionId uint) error {
	var transaction models.Transaction
	if err := initializers.DB.Preload("Splits").First(&transaction, transactionId).Error; err != nil {
		return err
	}
	if transaction.UserId != userId {
		return errors.New("transaction doesn't exist for this user")
	}
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, splitTratransaction := range transaction.Splits {
		if splitTratransaction.SettledTransactionId != nil {
			tx.Rollback()
			return errors.New("one of the splits is already settle,unsettle it first")
		}
		if err := tx.Delete(&splitTratransaction).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	transaction.SplitTag = "Will Split"
	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetSplitsOfTransaction(transaction models.Transaction, categoryIdToNameMap *map[uint]string, friendIdToNameMap *map[uint]string, unSettled bool, settled bool) ([]views.SplitView, error) {
	splits := make([]views.SplitView, 0)
	for _, splitTransaction := range transaction.Splits {
		if !settled {
			if splitTransaction.SettledTransactionId != nil {
				continue
			}
		}
		if !unSettled {
			if splitTransaction.SettledTransactionId == nil {
				continue
			}
		}
		categoryId := transaction.CategoryMappings[0].CategoryId
		categoryName := (*categoryIdToNameMap)[categoryId]
		friendName := (*friendIdToNameMap)[splitTransaction.FriendId]
		var settleTransactionId uint
		if splitTransaction.SettledTransactionId != nil {
			settleTransactionId = *splitTransaction.SettledTransactionId
		}
		split := views.SplitView{
			SplitTransactionId:   splitTransaction.ID,
			SourceTransactionId:  splitTransaction.SourceTransactionId,
			SettledTransactionId: settleTransactionId,
			CategoryName:         categoryName,
			SourceAmount:         transaction.Amount,
			Amount:               splitTransaction.Amount,
			FriendName:           friendName,
		}
		splits = append(splits, split)
	}
	return splits, nil
}

func GetSplitTransactions(userId uint, unSettled bool, settled bool) ([]views.SplitView, error) {
	var user models.User
	splits := make([]views.SplitView, 0)
	if err := initializers.DB.
		Preload("Transactions").
		Preload("Friends").
		Preload("Categories").
		Preload("Transactions.Splits").
		Preload("Transactions.CategoryMappings").
		First(&user, userId).
		Error; err != nil {
		return splits, err
	}
	categoryIdToNameMap := make(map[uint]string)
	for _, category := range user.Categories {
		categoryIdToNameMap[category.ID] = category.CategoryName
	}
	friendIdToNameMap := make(map[uint]string)
	for _, friend := range user.Friends {
		friendIdToNameMap[friend.ID] = friend.FriendName
	}
	for _, transaction := range user.Transactions {
		transactionSplits, err := GetSplitsOfTransaction(transaction, &categoryIdToNameMap, &friendIdToNameMap, unSettled, settled)
		if err != nil {
			return splits, err
		}
		splits = append(splits, transactionSplits...)
	}
	sort.Sort(views.BySplitTransactionIdDesc(splits))
	return splits, nil
}

func GetMoenyLentToFriendByCategory(userId uint, friendName string) (map[string][]views.SplitView, error) {
	moneyLentByCategory := make(map[string][]views.SplitView)
	splits, err := GetSplitTransactions(userId, true, false)
	if err != nil {
		return moneyLentByCategory, err
	}
	for _, split := range splits {
		if split.FriendName == friendName {
			if _, exists := moneyLentByCategory[split.CategoryName]; !exists {
				moneyLentByCategory[split.CategoryName] = make([]views.SplitView, 0)
			}
			moneyLentByCategory[split.CategoryName] = append(moneyLentByCategory[split.CategoryName], split)
		}
	}
	return moneyLentByCategory, nil
}

func SettleSplitsOfFriend(userId uint, friendId uint) error {
	var user models.User
	if err := initializers.DB.
		Preload("Transactions").
		Preload("Transactions.CategoryMappings").
		Preload("Friends").
		Preload("Friends.Debts").
		First(&user, userId).
		Error; err != nil {
		return err
	}
	transactionCategoryIdMap := make(map[uint]uint)
	for _, transaction := range user.Transactions {
		transactionCategoryIdMap[transaction.ID] = transaction.CategoryMappings[0].CategoryId
	}
	var friend models.Friend
	for _, friendIter := range user.Friends {
		if friendIter.ID == friendId {
			friend = friendIter
		}
	}
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, splitTransaction := range friend.Debts {
		if splitTransaction.SettledTransactionId == nil {
			categoryId := transactionCategoryIdMap[splitTransaction.SourceTransactionId]
			SettledTransactionId, err := CreateTransactionV2(userId, time.Now(), -1*splitTransaction.Amount, categoryId, "settle trans", "SETTLED TRANS")
			if err != nil {
				tx.Rollback()
				return err
			}
			splitTransaction.SettledTransactionId = &SettledTransactionId
			err = initializers.DB.Save(&splitTransaction).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
