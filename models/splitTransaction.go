package models

import (
	"gorm.io/gorm"
)

type SplitTransaction struct {
	gorm.Model
	SourceTransactionId  uint `gorm:"uniqueIndex:idx_transaction_friend_deleted;index;foreignkey:SourceTransactionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount               int
	FriendId             uint           `gorm:"uniqueIndex:idx_transaction_friend_deleted;index;foreignkey:FriendId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SettledTransactionId *uint          `gorm:"unique;foreignKey:SettledTransactionId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt            gorm.DeletedAt `gorm:"uniqueIndex:idx_transaction_friend_deleted;index"`
}
