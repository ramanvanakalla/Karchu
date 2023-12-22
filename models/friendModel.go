package models

import (
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	UserId     uint               `gorm:"uniqueIndex:idx_member"`
	FriendName string             `gorm:"uniqueIndex:idx_member"`
	Debts      []SplitTransaction `gorm:"foreignKey:FriendId"`
}
