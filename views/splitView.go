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

type BySplitTransactionIdDesc []SplitView

func (a BySplitTransactionIdDesc) Len() int      { return len(a) }
func (a BySplitTransactionIdDesc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySplitTransactionIdDesc) Less(i, j int) bool {
	return a[i].SplitTransactionId > a[j].SplitTransactionId
}

func (splitView SplitView) ToString() string {
	return fmt.Sprintf("Id: %d|SrcId: %d|SettId: %d|SrcAmnt: %d|Amount: %d|Friend: %s|Category: %s", splitView.SplitTransactionId, splitView.SourceTransactionId, splitView.SettledTransactionId, splitView.SourceAmount, splitView.Amount, splitView.FriendName, splitView.CategoryName)
}
