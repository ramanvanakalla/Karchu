package dao

import (
	"Karchu/initializers"
	"Karchu/models"
)

func CreateFriend(userId uint, friendName string) (uint, error) {
	newFriend := models.Friend{UserId: userId, FriendName: friendName}
	err := initializers.DB.Create(&newFriend).Error
	return newFriend.ID, err
}
