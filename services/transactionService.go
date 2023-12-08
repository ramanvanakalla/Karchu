package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/models"
	"fmt"
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

func transactionToString(transaction *models.Transaction) (string, *exceptions.GeneralException) {
	categoryName, err := dao.GetCategoryNameFromId(transaction.CategoryId)
	if err != nil {
		return "", exceptions.InternalServerError(err.Error(), "TRANSACTION_TO_STR_FAIL")
	}
	return fmt.Sprintf("Id: %d|Amount: %d|Category: %s|splitTag: %s|Desc: %s", transaction.ID, transaction.Amount, categoryName, transaction.SplitTag, transaction.Description), nil
}

func GetLastNTransactionsList(userId uint, lastN int) ([]string, *exceptions.GeneralException) {
	transactions, err := dao.GetLastNTransactionsByUserId(userId, lastN)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for _, transaction := range transactions {
		transStr, ex := transactionToString(&transaction)
		if ex != nil {
			return nil, ex
		}
		transactionsList = append(transactionsList, transStr)
	}
	return transactionsList, nil
}

func GetTransactionsList(userId uint) ([]string, *exceptions.GeneralException) {
	transactions, err := dao.GetAllTransactionsByUserId(userId)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for _, transaction := range transactions {
		transStr, ex := transactionToString(&transaction)
		if ex != nil {
			return nil, ex
		}
		transactionsList = append(transactionsList, transStr)
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
