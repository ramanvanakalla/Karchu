package models

import (
	"gorm.io/gorm"
)

type ModelSplit struct {
	gorm.Model
	ModelSplitId    uint           `gorm:"uniqueIndex:idx_friend_deleted;index"`
	FriendId        uint           `gorm:"uniqueIndex:idx_friend_deleted;index;foreignkey:FriendId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt       gorm.DeletedAt `gorm:"uniqueIndex:idx_friend_deleted;index"`
	SplitPercentage int
}
