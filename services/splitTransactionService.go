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
