package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/requests"
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

func SplitTransactionWithOneFriend(userId uint, transactionId uint, friendName string, amount int) *exceptions.GeneralException {
	splits := make([]requests.FriendSplit, 0)
	friendId, err := dao.GetFriendId(userId, friendName)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "FAIL_GETTING_FRND_ID")
	}
	splits = append(splits, requests.FriendSplit{FriendId: friendId, Amount: amount})
	return SplitTransaction(userId, transactionId, splits)
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

func UnSettleTransaction(userId uint, splitTransactionId uint) *exceptions.GeneralException {
	if err := dao.VerifySplitTransaction(userId, splitTransactionId); err != nil {
		return exceptions.InternalServerError(err.Error(), "SETTLE_VERIFICATION_FAIL")
	}
	if err := dao.UnSettleSplitTransaction(userId, splitTransactionId); err != nil {
		return exceptions.InternalServerError(err.Error(), "UNSETTLE_TRANS_FAIL")
	}
	return nil
}
