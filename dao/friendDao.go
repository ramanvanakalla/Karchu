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

func GetFriendId(userId uint, friendName string) (uint, error) {
	var friend models.Friend
	err := initializers.DB.Where(models.Friend{UserId: userId, FriendName: friendName}).First(&friend).Error
	return friend.ID, err
}

func GetFriendNameFromId(friendId uint) (string, error) {
	var friend models.Friend
	err := initializers.DB.
		First(&friend, friendId).
		Error
	return friend.FriendName, err
}

func GetFriends(userId uint) ([]string, error) {
	var user models.User
	friends := make([]string, 0)
	err := initializers.DB.
		Preload("Friends").
		First(&user, userId).
		Error
	if err != nil {
		return friends, err
	}
	for _, friend := range user.Friends {
		friends = append(friends, friend.FriendName)
	}
	return friends, nil
}

func IsFriends(userId uint, friendIds []uint) (bool, error) {
	var user models.User
	err := initializers.DB.Preload("Friends").First(&user, userId).Error
	if err != nil {
		return false, err
	}
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
