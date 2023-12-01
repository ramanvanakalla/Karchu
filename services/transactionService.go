package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/models"
	"fmt"
	"time"
)

func CreateTransaction(userId uint, time time.Time, amount int, category string, description string, splitTag string) (uint, *exceptions.GeneralException) {
	categoryId, err := dao.GetCategoryIdByUserIdAndCategoryName(userId, category)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "CANT_GET_CATEGORY")
	}
	transactionId, err := dao.CreateTransaction(userId, time, amount, category, categoryId, description, splitTag)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return transactionId, nil
}

func DeleteTransaction(transactionId uint, userId uint) (uint, *exceptions.GeneralException) {
	delTransactionId, err := dao.DeleteTransactionbyTransactionIdAndUserId(transactionId, userId)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DELETE_TRANS_FAIL")
	}
	return delTransactionId, nil
}

func DeleteTransactionFromTransString(TransString string, userId uint) (uint, *exceptions.GeneralException) {
	transaction, err := StringToTransaction(TransString)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "CANT_PARSE_TRANS_STRING")
	}
	return DeleteTransaction(transaction.ID, userId)
}

func StringToTransaction(input string) (*models.Transaction, error) {
	var transaction models.Transaction
	fmt.Println(input)
	_, err := fmt.Sscanf(input, "Id: %d|Amount: %d|category: %[^\n]|splitTag: %[^\n]|Desc: %[^\n]",
		&transaction.ID, &transaction.Amount, &transaction.Category, &transaction.SplitTag, &transaction.Description)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func transactionToString(transaction *models.Transaction) string {
	return fmt.Sprintf("Id: %d|Amount: %d|category: %s|splitTag: %s|Desc: %s", transaction.ID, transaction.Amount, transaction.Category, transaction.SplitTag, transaction.Description)
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

func GetNetMoneySpentByCategory(userId uint) ([]string, *exceptions.GeneralException) {
	categoriesAndSum, err := dao.GetNetMoneySpentByCategory(userId)
	netByCategoriesList := make([]string, 0)
	if err != nil {
		return netByCategoriesList, exceptions.InternalServerError(err.Error(), "NET_CATEGORY_SUM_GET_FAIL")
	}
	for _, categorySum := range categoriesAndSum {
		netByCategoriesList = append(netByCategoriesList, categorySum.ToString())
	}
	return netByCategoriesList, nil
}
