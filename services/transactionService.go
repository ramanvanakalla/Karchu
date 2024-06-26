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

func CreateTransactionAndModelSplit(userId uint, amount int, category string, description string, modelSplit string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&category) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", category), "INVALID_CATEGORY_FORMAT")
	}
	categoryId, err := dao.GetCategoryIdByUserIdAndCategoryName(userId, category)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "CANT_GET_CATEGORY")
	}
	modelSplitMapId, err := dao.GetModelSplitMappingId(userId, modelSplit)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "CANT_GET_MODEL_SPLIT")
	}
	fmt.Printf("model split id %d", modelSplitMapId)
	transactionId, err := dao.CreateTransactionAndModelSplit(userId, time.Now(), amount, categoryId, description, modelSplitMapId)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "FAILED_TO_ADD_TRANSACTION")
	}
	return transactionId, nil
}

func CreateTransactionV2(userId uint, amount int, category string, description string, splitTag string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeCategory(&category) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", category), "INVALID_CATEGORY_FORMAT")
	}
	categoryId, err := dao.GetCategoryIdByUserIdAndCategoryName(userId, category)
	if err != nil {
		return 0, exceptions.BadRequestError(err.Error(), "CANT_GET_CATEGORY")
	}
	transactionId, err := dao.CreateTransactionV2(userId, time.Now(), amount, categoryId, description, splitTag)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return transactionId, nil
}

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

func GetLastNTransactionsListV2(userId uint, lastN int) ([]string, *exceptions.GeneralException) {
	transactions, err := dao.GetLastNTransactionsByUserIdV2(userId, lastN)
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

func GetTransactionsListV2(userId uint) ([]string, *exceptions.GeneralException) {
	transactionViewList, err := dao.GetTransactionsByUserId(userId)
	transactionsList := make([]string, 0)
	if err != nil {
		return transactionsList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	for _, transactionView := range transactionViewList {
		transactionsList = append(transactionsList, transactionView.ToString())
	}
	return transactionsList, nil
}

func GetTransactionsV2(userId uint) ([]views.TransactionWithCategory, *exceptions.GeneralException) {
	transactionViewList, err := dao.GetTransactionsByUserId(userId)
	if err != nil {
		return transactionViewList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	return transactionViewList, nil
}

func GetTransactionsFiltered(userId uint, startDate time.Time, endDate time.Time, categories []string, splitTag string) ([]views.TransactionWithCategory, *exceptions.GeneralException) {
	transactionViewList, err := dao.GetTransactionsByUserIdFiltered(userId, startDate, endDate, categories, splitTag)
	if err != nil {
		return transactionViewList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	return transactionViewList, nil
}

func GetNetMoneySpentByCategory(userId uint) ([]views.NetCategorySum, *exceptions.GeneralException) {
	allTransactions, err := dao.GetTransactionsByUserId(userId)
	netCategorySumList := make([]views.NetCategorySum, 0)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	netCategorySumMap := make(map[string]int)
	for _, transactionView := range allTransactions {
		if _, exists := netCategorySumMap[transactionView.CategoryName]; !exists {
			netCategorySumMap[transactionView.CategoryName] = 0
		}
		netCategorySumMap[transactionView.CategoryName] += transactionView.Amount
	}
	for category, netAmount := range netCategorySumMap {
		netCategorySum := views.NetCategorySum{Category: category, NetAmount: netAmount}
		netCategorySumList = append(netCategorySumList, netCategorySum)
	}
	allCategories, err := dao.GetCategoriesByUserID(userId)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}

	for _, category := range allCategories {
		_, exists := netCategorySumMap[category.CategoryName]
		if !exists {
			netCategorySum := views.NetCategorySum{Category: category.CategoryName, NetAmount: 0}
			netCategorySumList = append(netCategorySumList, netCategorySum)
		}
	}

	sort.Sort(views.ByNetAmountDesc(netCategorySumList))
	return netCategorySumList, nil
}

func GetNetMoneySpentByCategoryFiltered(userId uint, startDate time.Time, endDate time.Time) ([]views.NetCategorySum, *exceptions.GeneralException) {
	categories := make([]string, 0)
	allTransactions, err := dao.GetTransactionsByUserIdFiltered(userId, startDate, endDate, categories, "")
	netCategorySumList := make([]views.NetCategorySum, 0)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	netCategorySumMap := make(map[string]int)
	for _, transactionView := range allTransactions {
		if _, exists := netCategorySumMap[transactionView.CategoryName]; !exists {
			netCategorySumMap[transactionView.CategoryName] = 0
		}
		netCategorySumMap[transactionView.CategoryName] += transactionView.Amount
	}
	for category, netAmount := range netCategorySumMap {
		netCategorySum := views.NetCategorySum{Category: category, NetAmount: netAmount}
		netCategorySumList = append(netCategorySumList, netCategorySum)
	}
	allCategories, err := dao.GetCategoriesByUserID(userId)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}

	for _, category := range allCategories {
		_, exists := netCategorySumMap[category.CategoryName]
		if !exists {
			netCategorySum := views.NetCategorySum{Category: category.CategoryName, NetAmount: 0}
			netCategorySumList = append(netCategorySumList, netCategorySum)
		}
	}

	sort.Sort(views.ByNetAmountDesc(netCategorySumList))
	return netCategorySumList, nil
}

func GetNetMoneySpentByCategory2(userId uint) ([]string, *exceptions.GeneralException) {
	allTransactions, err := dao.GetTransactionsByUserId(userId)
	netCategorySumList := make([]string, 0)
	if err != nil {
		return netCategorySumList, exceptions.InternalServerError(err.Error(), "TRANSACTION_GET_FAIL")
	}
	totalMoneySpent := 0
	netCategorySumMap := make(map[string]int)
	for _, transactionView := range allTransactions {
		if _, exists := netCategorySumMap[transactionView.CategoryName]; !exists {
			netCategorySumMap[transactionView.CategoryName] = 0
		}
		netCategorySumMap[transactionView.CategoryName] += transactionView.Amount
	}
	for category, netAmount := range netCategorySumMap {
		totalMoneySpent += netAmount
		netCategorySum := views.NetCategorySum{Category: category, NetAmount: netAmount}
		netCategorySumList = append(netCategorySumList, netCategorySum.ToString())
	}
	totalMoney := views.NetCategorySum{Category: "Total Money Spent", NetAmount: totalMoneySpent}
	netCategorySumList = append(netCategorySumList, totalMoney.ToString())
	return netCategorySumList, nil
}

func CreateTransactionAndSplitWithOne(userId uint, time time.Time, amount int, category string, description string, splitTag string, friendName string, splitAmount int) *exceptions.GeneralException {
	fmt.Println(description)
	if !validateAndNormalizeCategory(&category) {
		return exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", category), "INVALID_CATEGORY_FORMAT")
	}
	if splitAmount > amount {
		return exceptions.BadRequestError("split amount can not be greater than trans amount", "SPLT_GREATER_THAN_TRANS_AMT")
	}
	categoryId, err := dao.GetCategoryIdByUserIdAndCategoryName(userId, category)
	if err != nil {
		return exceptions.BadRequestError(err.Error(), "CANT_GET_CATEGORY")
	}
	friendId, err := dao.GetFriendId(userId, friendName)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_GETTING_FRND_ID")
	}
	if err := dao.CreateTransactionAndSplitWithOne(userId, time, amount, categoryId, description, splitTag, friendId, splitAmount); err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_TRANS_ADD_SPLIT")
	}
	return nil
}
