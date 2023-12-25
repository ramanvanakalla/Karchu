package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/requests"
	"Karchu/views"
	"errors"
)

func verifySplits(userId uint, transactionId uint, Splits []requests.FriendSplit) error {
	friendIds := make([]uint, 0)
	netAmount := 0
	for _, split := range Splits {
		friendIds = append(friendIds, uint(split.FriendId))
		netAmount += split.Amount
	}
	isFriends, err := dao.IsFriends(userId, friendIds)
	if err != nil {
		return err
	}
	if !isFriends {
		return errors.New("ids are not friends to user")
	}
	transaction, err := dao.GetTransaction(transactionId)
	if err != nil {
		return err
	}
	if netAmount > transaction.Amount {
		return errors.New("given splits sum up to more than amount of transaction")
	}
	return nil
}

func SplitTransactionWithOneFriend(userId uint, TransString string, friendName string, amount int) *exceptions.GeneralException {
	splits := make([]requests.FriendSplit, 0)
	transaction, err := StringToTransaction(TransString)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_GETTING_TRANS_ID")
	}
	friendId, err := dao.GetFriendId(userId, friendName)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_GETTING_FRND_ID")
	}
	splits = append(splits, requests.FriendSplit{FriendId: friendId, Amount: amount})
	return SplitTransaction(userId, transaction.ID, splits)
}

func GetSplitTransactions(userId uint) ([]string, *exceptions.GeneralException) {
	splits, err := dao.GetSplitTransactions(userId, true, true)
	splitStrings := make([]string, 0)
	if err != nil {
		return splitStrings, exceptions.InternalServerError(err.Error(), "FAIL_GETTING_SPLITS")
	}
	for _, split := range splits {
		splitStrings = append(splitStrings, split.ToString())
	}
	return splitStrings, nil
}

func GetUnSettledSplitTransactions(userId uint) ([]string, *exceptions.GeneralException) {
	splits, err := dao.GetSplitTransactions(userId, true, false)
	splitStrings := make([]string, 0)
	if err != nil {
		return splitStrings, exceptions.InternalServerError(err.Error(), "FAIL_GETTING_SPLITS")
	}
	for _, split := range splits {
		splitStrings = append(splitStrings, split.ToString())
	}
	return splitStrings, nil
}

func GetSettledSplitTransactions(userId uint) ([]string, *exceptions.GeneralException) {
	splits, err := dao.GetSplitTransactions(userId, false, true)
	splitStrings := make([]string, 0)
	if err != nil {
		return splitStrings, exceptions.InternalServerError(err.Error(), "FAIL_GETTING_SPLITS")
	}
	for _, split := range splits {
		splitStrings = append(splitStrings, split.ToString())
	}
	return splitStrings, nil
}

func SplitTransaction(userId uint, transactionId uint, splits []requests.FriendSplit) *exceptions.GeneralException {
	err := verifySplits(userId, transactionId, splits)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "SPLIT_TRANSACTION_FAIL")
	}
	alreadySplit, err := dao.TransactionAlreadySplit(transactionId)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "SPLIT_TRANSACTION_FAIL")
	}
	if alreadySplit {
		return exceptions.InternalServerError("Transaction is already split", "SPLIT_TRANSACTION_FAIL")
	}
	err = dao.AddSplitTransactions(userId, transactionId, splits)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "SPLIT_TRANSACTION_FAIL")
	}
	return nil
}

func DeleteSplitTransactionString(userId uint, transString string) *exceptions.GeneralException {
	transaction, err := StringToTransaction(transString)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_GETTING_TRANS_ID")
	}
	alreadySplit, err := dao.TransactionAlreadySplit(transaction.ID)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "DELETE_SPLIT_TRANS_FAIL")
	}
	if !alreadySplit {
		return exceptions.InternalServerError("Transaction is not split", "TRANSACTION_NOT_SPLIT")
	}
	if err := dao.DeleteTransactionSplit(userId, transaction.ID); err != nil {
		return exceptions.InternalServerError(err.Error(), "DELETE_SPLIT_TRANS_FAIL")
	}
	return nil
}

func DeleteSplitTransaction(userId uint, transactionId uint) *exceptions.GeneralException {
	alreadySplit, err := dao.TransactionAlreadySplit(transactionId)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "DELETE_SPLIT_TRANS_FAIL")
	}
	if !alreadySplit {
		return exceptions.InternalServerError("Transaction is not split", "TRANSACTION_NOT_SPLIT")
	}
	if err := dao.DeleteTransactionSplit(userId, transactionId); err != nil {
		return exceptions.InternalServerError(err.Error(), "DELETE_SPLIT_TRANS_FAIL")
	}
	return nil
}

func SettleTransaction(userId uint, splitTransactionId uint) *exceptions.GeneralException {
	err := dao.VerifySplitTransaction(userId, splitTransactionId)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "SETTLE_VERIFICATION_FAIL")
	}
	err = dao.SettleTransaction(userId, splitTransactionId)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "SETTLEMENT_FAILED")
	}
	return nil
}

func SettleTransactionString(userId uint, splitTransactionString string) *exceptions.GeneralException {
	splitTransaction, err := StringToTransaction(splitTransactionString)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "TRANS_TO_STR_FAIL")
	}
	return SettleTransaction(userId, splitTransaction.ID)
}

func UnSettleTransaction(userId uint, splitTransactionId uint) *exceptions.GeneralException {
	if err := dao.VerifySplitTransaction(userId, splitTransactionId); err != nil {
		return exceptions.InternalServerError(err.Error(), "SETTLE_VERIFICATION_FAIL")
	}
	if err := dao.UnSettleSplitTransaction(userId, splitTransactionId); err != nil {
		return exceptions.InternalServerError(err.Error(), "UNSETTLE_TRANS_FAIL")
	}
	return nil
}

func SettleSplitsOfFriend(userId uint, friendName string) *exceptions.GeneralException {
	friendId, err := dao.GetFriendId(userId, friendName)
	if err != nil {
		exceptions.InternalServerError(err.Error(), "GET_FRIEND_FAIL")
	}
	err = dao.SettleSplitsOfFriend(userId, friendId)
	if err != nil {
		exceptions.InternalServerError(err.Error(), "SETTLE_FAIL")
	}
	return nil
}

func UnSettleTransactionString(userId uint, splitTransactionString string) *exceptions.GeneralException {
	splitTransaction, err := StringToTransaction(splitTransactionString)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "TRANS_TO_STR_FAIL")
	}
	return UnSettleTransaction(userId, splitTransaction.ID)
}

func MoneyLentToFriend(userId uint, friendName string) ([]string, *exceptions.GeneralException) {
	categoryLevelMoneyLent := make([]string, 0)
	SplitsByCategory, err := dao.GetMoenyLentToFriendByCategory(userId, friendName)
	if err != nil {
		return categoryLevelMoneyLent, exceptions.InternalServerError(err.Error(), "FAIL_GETTING_SPLITS")
	}
	totalMoneyLent := 0
	for categoryName, splits := range SplitsByCategory {
		netAmount := 0
		for _, split := range splits {
			netAmount += split.Amount
		}
		categorySum := views.NetCategorySum{Category: categoryName, NetAmount: netAmount}
		categoryLevelMoneyLent = append(categoryLevelMoneyLent, categorySum.ToString())
		totalMoneyLent += netAmount
	}
	categorySum := views.NetCategorySum{Category: "Total Money Lent", NetAmount: totalMoneyLent}
	categoryLevelMoneyLent = append(categoryLevelMoneyLent, categorySum.ToString())
	return categoryLevelMoneyLent, nil
}
