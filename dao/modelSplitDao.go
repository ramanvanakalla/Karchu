package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"Karchu/requests"
	"fmt"
)

func CreateModelSplitForUser(userId uint, modelSplitName string, friendSplitPercentageSplits []requests.FriendSplitPercentage) error {
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// step1: add a row for the current user with new model split
	modelSplitMap := models.ModelSplitMapping{UserId: userId, ModelSplitName: modelSplitName}
	modelSplitMapErr := tx.Create(&modelSplitMap).Error
	if modelSplitMapErr != nil {
		tx.Rollback()
		return modelSplitMapErr
	}
	fmt.Println(modelSplitMap.ID)
	// step2: take the model mapping id and add percentags for all the friendids
	var modelSplits []models.ModelSplit
	for _, friendSplitPercentageSplit := range friendSplitPercentageSplits {
		modelSplits = append(modelSplits, models.ModelSplit{
			ModelSplitId:    modelSplitMap.ID,
			FriendId:        friendSplitPercentageSplit.FriendId,
			SplitPercentage: friendSplitPercentageSplit.SplitPercentage,
		})
	}

	// step2: Bulk insert model splits
	if err := tx.Create(&modelSplits).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetModelSplits(userId uint) ([]string, error) {
	modelSplitNames := make([]string, 0)
	modelSplits := make([]models.ModelSplitMapping, 0)
	err := initializers.DB.Where(models.ModelSplitMapping{UserId: userId}).Find(&modelSplits).Error
	if err != nil {
		return modelSplitNames, err
	}
	for _, modelSplit := range modelSplits {
		modelSplitNames = append(modelSplitNames, modelSplit.ModelSplitName)
	}
	return modelSplitNames, nil
}
