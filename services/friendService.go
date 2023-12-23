package services

import (
	"Karchu/dao"
	"Karchu/exceptions"
	"fmt"
	"strings"
)

func validateAndNormalizeFriendName(friendName *string) bool {
	*friendName = strings.TrimSpace(*friendName)
	if len(*friendName) < 1 {
		return false
	} else {
		return true
	}
}

func CreateFriend(userId uint, friendName string) (uint, *exceptions.GeneralException) {
	if !validateAndNormalizeFriendName(&friendName) {
		return 0, exceptions.BadRequestError(fmt.Sprintf("invalid category format %s", friendName), "INVALID_FRIEND_FORMAT")
	}
	friendId, err := dao.CreateFriend(userId, friendName)
	if err != nil {
		return 0, exceptions.InternalServerError(err.Error(), "DB_INSERTION_FAIL")
	}
	return friendId, nil
}

func GetFriends(userId uint) ([]string, *exceptions.GeneralException) {
	friends, err := dao.GetFriends(userId)
	if err != nil {
		return friends, exceptions.InternalServerError(err.Error(), "GET_FRIENDS_FAIL")
	}
	return friends, nil
}
