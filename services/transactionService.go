package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/models"
	"Karchu/views"
	"fmt"
	"sort"
	"strings"
	"time"
)

func CreateTransaction(userId uint, time time.Time, amount int, category string, description string, splitTag string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&category) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", category), "INVALID_CATEGORY_FORMAT")
	}
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
	fields := strings.Split(input, "|")
	var transaction models.Transaction
	for _, field := range fields {
		keyValue := strings.SplitN(field, ":", 2)
		if len(keyValue) == 2 {
			key, value := strings.TrimSpace(keyValue[0]), strings.TrimSpace(keyValue[1])
			switch key {
			case "Id":
				fmt.Sscanf(value, "%d", &transaction.ID)
			case "Amount":
				fmt.Sscanf(value, "%d", &transaction.Amount)
			case "category":
			case "splitTag":
				transaction.SplitTag = value
			case "Desc":
				transaction.Description = value
			}
		}
	}

	return &transaction, nil
}

func GetLastNTransactionsList(userId uint, lastN int) ([]string, *exceptions.GeneralException) {
	transactions, err := dao.GetLastNTransactionsByUserId(userId, lastN)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for _, transaction := range transactions {
		transStr := transaction.ToString()
		transactionsList = append(transactionsList, transStr)
	}
	return transactionsList, nil
}

func GetTransactionsList(userId uint) ([]string, *exceptions.GeneralException) {
	categoryTransactionsMap, err := dao.GetAllTransactionsByUserId(userId)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	transactionViewList := make([]views.TransactionWithCategory, 0)
	for categoryName, transactions := range categoryTransactionsMap {
		for _, transaction := range transactions {
			transactionViewList = append(transactionViewList, views.NewTransactionWithCategory(transaction, categoryName))
		}
	}
	sort.Slice(transactionViewList, func(i, j int) bool {
		return transactionViewList[i].ID > transactionViewList[j].ID
	})
	for _, transactionView := range transactionViewList {
		transactionsList = append(transactionsList, transactionView.ToString())
	}
	return transactionsList, nil
}

func GetTransactions(userId uint) ([]views.TransactionWithCategory, *exceptions.GeneralException) {
	categoryTransactionsMap, err := dao.GetAllTransactionsByUserId(userId)
	transactionsList := make([]views.TransactionWithCategory, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for categoryName, transactions := range categoryTransactionsMap {
		for _, transaction := range transactions {
			transactionView := views.NewTransactionWithCategory(transaction, categoryName)
			transactionsList = append(transactionsList, transactionView)
		}
	}
	return transactionsList, nil
}

func GetNetMoneySpentByCategory(userId uint) ([]string, *exceptions.GeneralException) {
	categoryTransactionsMap, err := dao.GetAllTransactionsByUserId(userId)
	netCategorySumList := make([]string, 0)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for categoryName, transactions := range categoryTransactionsMap {
		netAmount := 0
		for _, transaction := range transactions {
			netAmount += transaction.Amount
		}
		netCategorySum := views.NetCategorySum{Category: categoryName, NetAmount: netAmount}
		netCategorySumList = append(netCategorySumList, netCategorySum.ToString())
	}
	return netCategorySumList, nil
}
