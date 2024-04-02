package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"Karchu/requests"
	"errors"
)

func verifyModelSplits(userId uint, friendSplits []requests.FriendSplitPercentage) error {
	friendIds := make([]uint, 0)
	totalPercentage := 0
	for _, ModelSplit := range friendSplits {
		friendIds = append(friendIds, uint(ModelSplit.FriendId))
		totalPercentage += ModelSplit.SplitPercentage
	}
	isFriends, err := dao.IsFriends(userId, friendIds)
	if err != nil {
		return err
	}
	if !isFriends {
		return errors.New("ids are not friends to user")
	}
	if totalPercentage > 100 {
		return errors.New("sum of percentages can not be greater than 100")
	}
	return nil
}

func CreateModelSplit(userId uint, modelSplitName string, friendSplits []requests.FriendSplitPercentage) *exceptions.GeneralException {
	err := verifyModelSplits(userId, friendSplits)
	if err != nil {
		return exceptions.InternalServerError(err.Error(), "MODEL_SPLIT_FAIL")
	}
	dbErr := dao.CreateModelSplitForUser(userId, modelSplitName, friendSplits)
	if dbErr != nil {
		return exceptions.InternalServerError(dbErr.Error(), "DB_INSERTION_FAIL")
	}
	return nil
}

func GetModelSplits(userId uint) ([]string, *exceptions.GeneralException) {
	modelSplits, err := dao.GetModelSplitNames(userId)
	if err != nil {
		return modelSplits, exceptions.InternalServerError(err.Error(), "MODEL_SPLIT_GET_FAIL")
	}
	return modelSplits, nil
}
