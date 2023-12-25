package views

import "fmt"

type SplitView struct {
	SourceTransactionId  uint
	SplitTransactionId   uint
	SettledTransactionId uint
	CategoryName         string
	SourceAmount         int
	Amount               int
	FriendName           string
}

func (splitView SplitView) ToString() string {
	return fmt.Sprintf("Id: %d|SrcId: %d|SettId: %d|SrcAmnt: %d|Amount: %d|Friend: %s|Category: %s", splitView.SplitTransactionId, splitView.SourceTransactionId, splitView.SettledTransactionId, splitView.SourceAmount, splitView.Amount, splitView.FriendName, splitView.CategoryName)
}
