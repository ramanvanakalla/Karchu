package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/requests"
	"Karchu/views"
	"errors"
	"fmt"
	"time"
)

func CreateTransactionAndModelSplit(userId uint, time time.Time, amount int, categoryId uint, description string, modelSplitMapId uint) (uint, error) {
	transactionId, err := CreateTransactionV2(userId, time, amount, categoryId, description, "Model Split")
	fmt.Printf("%d is created\n", transactionId)
	if err != nil {
		return transactionId, err
	}
	friendsToShareMap, err := GetModelSplitsOfId(modelSplitMapId)
	fmt.Println(friendsToShareMap)
	if err != nil {
		return transactionId, err
	}
	splits := make([]requests.FriendSplit, 0)
	for friendId, percentageShare := range friendsToShareMap {
		splits = append(splits, requests.FriendSplit{FriendId: friendId, Amount: amount * percentageShare / 100})
	}
	err = AddSplitTransactions(userId, transactionId, splits)
	if err != nil {
		return transactionId, err
	}
	return transactionId, nil
}

func CreateTransactionV2(userId uint, time time.Time, amount int, categoryId uint, description string, splitTag string) (uint, error) {
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// step 1: add transaction to table
	transaction := models.Transaction{UserId: userId, Time: time, Amount: amount, Description: description, SplitTag: splitTag}
	transactionErr := tx.Create(&transaction).Error
	if transactionErr != nil {
		tx.Rollback()
		return 0, transactionErr
	}
	// step 2: add to transactionCatgoryTable
	categoryTransactionMap := models.CategoryTransactionMapping{CategoryId: categoryId, TransactionId: transaction.ID}
	categoryTransactionErr := tx.Create(&categoryTransactionMap).Error
	if categoryTransactionErr != nil {
		tx.Rollback()
		return 0, categoryTransactionErr
	}
	return transaction.ID, tx.Commit().Error
}

func CreateTransactionAndSplitWithOne(userId uint, time time.Time, amount int, categoryId uint, description string, splitTag string, friendId uint, splitAmount int) error {
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// step 1: add transaction to table
	transaction := models.Transaction{UserId: userId, Time: time, Amount: amount, Description: description, SplitTag: splitTag}
	transactionErr := tx.Create(&transaction).Error
	if transactionErr != nil {
		tx.Rollback()
		return transactionErr
	}
	// step 2: add to transactionCatgoryTable
	categoryTransactionMap := models.CategoryTransactionMapping{CategoryId: categoryId, TransactionId: transaction.ID}
	categoryTransactionErr := tx.Create(&categoryTransactionMap).Error
	if categoryTransactionErr != nil {
		tx.Rollback()
		return categoryTransactionErr
	}
	splitTransaction := models.SplitTransaction{SourceTransactionId: transaction.ID, Amount: splitAmount, FriendId: friendId}
	err := tx.Create(&splitTransaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func CreateTransaction(userId uint, time time.Time, amount int, category string, categoryId uint, description string, splitTag string) (uint, error) {
	transaction := models.Transaction{UserId: userId, Time: time, Amount: amount, Description: description, SplitTag: splitTag}
	err := initializers.DB.Create(&transaction).Error
	return transaction.ID, err
}

func GetTransaction(transactionId uint) (models.Transaction, error) {
	var transaction models.Transaction
	err := initializers.DB.First(&transaction, transactionId).Error
	return transaction, err
}

func IsTransactionSettledTrans(transactionId uint) (bool, error) {
	var transaction models.Transaction
	if err := initializers.DB.Preload("SourceSplitTransaction").First(&transaction, transactionId).Error; err != nil {
		return true, err
	}
	if transaction.SourceSplitTransaction.ID == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func IsTransactionSplit(transactionId uint) (bool, error) {
	var transaction models.Transaction
	if err := initializers.DB.Preload("Splits").First(&transaction, transactionId).Error; err != nil {
		return true, err
	}
	if len(transaction.Splits) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func DeleteTransactionbyTransactionIdAndUserId(transactionId uint, userId uint) (uint, error) {
	settledTransaction, err := IsTransactionSettledTrans(transactionId)
	if err != nil {
		return 0, err
	}
	if settledTransaction {
		return 0, errors.New("settled transaction can not be deleted, it has to be unsettled")
	}

	isTransactionSplit, err := IsTransactionSplit(transactionId)
	if err != nil {
		return 0, err
	}
	if isTransactionSplit {
		return 0, errors.New("transaction has split, please delete the slit of this trans first")
	}
	var transaction models.Transaction
	err = initializers.DB.
		Model(&models.Transaction{}).
		Where("id = ? and user_id = ?", transactionId, userId).
		First(&transaction).
		Error
	if err != nil {
		return 0, err
	}
	deletionErr := initializers.DB.Delete(&transaction).Error
	return transaction.ID, deletionErr
}

func GetLastNTransactionsByUserIdV2(userId uint, lastN int) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	err := initializers.DB.
		Model(&models.Transaction{}).
		Joins("JOIN category_transaction_mappings ON transactions.id = category_transaction_mappings.transaction_id").
		Joins("JOIN categories ON category_transaction_mappings.category_id = categories.id").
		Where("transactions.user_id = ?", userId).
		Select("transactions.*,categories.category_name").
		Order("transactions.id desc").
		Limit(lastN).
		Find(&transactions).
		Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetLastNTransactionsByUserId(userId uint, lastN int) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	err := initializers.DB.
		Model(&models.Transaction{}).
		Select("transactions.*,categories.category_name").
		Joins("JOIN categories ON transactions.category_id = categories.id").
		Where("transactions.user_id = ?", userId).
		Order("transactions.created_at desc").
		Limit(lastN).
		Find(&transactions).
		Error
	return transactions, err
}

func GetTransactionsByUserId(userId uint) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	err := initializers.DB.
		Model(&models.Transaction{}).
		Joins("JOIN category_transaction_mappings ON transactions.id = category_transaction_mappings.transaction_id").
		Joins("JOIN categories ON category_transaction_mappings.category_id = categories.id AND categories.deleted_at IS NULL").
		Where("transactions.user_id = ?", userId).
		Select("transactions.*,categories.category_name").
		Order("transactions.id desc").
		Find(&transactions).
		Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransactionsByUserIdFiltered(userId uint, startDate time.Time, endDate time.Time, categories []string, splitTag string) ([]views.TransactionWithCategory, error) {
	var transactions []views.TransactionWithCategory
	query := initializers.DB.
		Model(&models.Transaction{}).
		Joins("JOIN category_transaction_mappings ON transactions.id = category_transaction_mappings.transaction_id").
		Joins("JOIN categories ON category_transaction_mappings.category_id = categories.id AND categories.deleted_at IS NULL").
		Where("transactions.user_id = ?", userId).
		Select("transactions.*,categories.category_name").
		Where("transactions.time >= ?", startDate).
		Where("transactions.time <= ?", endDate)
	if splitTag != "" {
		query = query.Where("transactions.split_tag = ?", splitTag)
	}
	if len(categories) > 0 {
		query = query.Where("categories.category_name in ?", categories)
	}
	err := query.
		Order("transactions.id desc").
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetNetMoneySpentByCategory(userID uint) ([]views.NetCategorySum, error) {
	var amountByCategory []views.NetCategorySum
	err := initializers.DB.
		Model(&models.Transaction{}).
		Select("category, sum(amount) as net_amount").
		Where("user_id = ?", userID).
		Group("category").
		Order("net_amount desc").
		Scan(&amountByCategory).
		Error
	return amountByCategory, err
}
