package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/models"
	"fmt"
	"time"
)

func CreateTransaction(userId uint, time time.Time, amount int, category string, description string, splitTag string, mapUrl string) (uint, *exceptions.GeneralException) {
	categoryId, err := dao.GetCategoryIdByUserIdAndCategoryName(userId, category)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "CANT_GET_CATEGORY")
	}
	transactionId, err := dao.CreateTransaction(userId, time, amount, category, categoryId, description, splitTag, mapUrl)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return transactionId, nil
}

func transactionToString(transaction *models.Transaction) string {
	return fmt.Sprintf("Amount: %d|category: %s|splitTag: %s|Desc: %s", transaction.Amount, transaction.Category, transaction.SplitTag, transaction.Description)
}

func GetLastNTransactionsList(userId uint, lastN int) ([]string, *exceptions.GeneralException) {
	transactions, err := dao.GetLastNTransactionsByUserId(userId, lastN)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for _, transaction := range transactions {
		transactionsList = append(transactionsList, transactionToString(&transaction))
	}
	return transactionsList, nil
}
