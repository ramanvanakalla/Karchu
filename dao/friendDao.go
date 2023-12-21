package dao

import (
	"Karchu/initializers"
	"Karchu/models"
	"fmt"
)

func CreateFriend(userId uint, friendName string) (uint, error) {
	newFriend := models.Friend{UserId: userId, FriendName: friendName}
	err := initializers.DB.Create(&newFriend).Error
	return newFriend.ID, err
}

func IsFriends(userId uint, friendIds []uint) (bool, error) {
	var user models.User
	err := initializers.DB.Preload("Friends").First(&user, userId).Error
	if err != nil {
		return false, err
	}
	fmt.Println(user)
	fmt.Println(user.Friends)
	idMap := make(map[uint]bool)
	for _, friend := range user.Friends {
		idMap[friend.ID] = true
	}
	for _, id := range friendIds {
		_, exists := idMap[id]
		if !exists {
			return false, nil
		}
	}
	return true, nil
}
