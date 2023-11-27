package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserId      uint
	CategoryId  uint
	Amount      int
	Time        time.Time
	Category    string
	Description string
	SplitTag    string
	MapUrl      string
}
