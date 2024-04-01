package requests

type UserReq struct {
	Email    string
	Password string
}

type CreateUserReq struct {
	Email    string
	Password string
	Name     string
}
type CreateCategoryReq struct {
	Email        string
	Password     string
	CategoryName string
}

type DeleteCategoryReq struct {
	Email        string
	Password     string
	CategoryName string
}

type TransactionsOfCategoryReq struct {
	Email        string
	Password     string
	CategoryName string
}

type CreateTransactionReq struct {
	Email       string
	Password    string
	Amount      int
	Category    string
	Description string
	SplitTag    string
}

type GetLastNTransactionsReq struct {
	Email    string
	Password string
	LastN    int
}

type GetTransactionsReq struct {
	Email    string
	Password string
}
type GetFilteredTransactionsReq struct {
	Email      string
	Password   string
	StartDate  string
	EndDate    string
	Categories []string
	SplitTag   string
}

type DeleteTransactionReq struct {
	Email         string
	Password      string
	TransactionId uint
}

type DeleteTransactionFromTransStringReq struct {
	Email       string
	Password    string
	TransString string
}

type NetAmountByCategoryReq struct {
	Email    string
	Password string
}

type NetAmountByCategoryFilteredReq struct {
	Email     string
	Password  string
	StartDate string
	EndDate   string
}

type RenameCategoryReq struct {
	Email           string
	Password        string
	OldCategoryName string
	NewCategoryName string
}

type MergeCategory struct {
	Email                   string
	Password                string
	SourceCategoryName      string
	DestinationCategoryName string
}

type CreateFriendReq struct {
	Email      string
	Password   string
	FriendName string
}

type SplitTransactionReq struct {
	Email         string
	Password      string
	TransactionId uint
	Splits        []FriendSplit
}

type ModelSplitTransactionReq struct {
	Email          string
	Password       string
	ModelSplitName string
	ModelSplits    []FriendSplitPercentage
}
type GetModelSplitReq struct {
	Email    string
	Password string
}

type FriendSplitPercentage struct {
	FriendId        uint
	SplitPercentage int
}

type FriendSplit struct {
	FriendId uint
	Amount   int
}

type SplitWithOneFriendReq struct {
	Email       string
	Password    string
	TransString string
	Friend      string
	Amount      int
}

type SettleTransactionReq struct {
	Email              string
	Password           string
	SplitTransactionId uint
}

type SettleTransactionStringReq struct {
	Email                  string
	Password               string
	SplitTransactionString string
}

type UnSettleTransactionReq struct {
	Email              string
	Password           string
	SplitTransactionId uint
}

type UnSettleTransactionStringReq struct {
	Email                  string
	Password               string
	SplitTransactionString string
}

type SettleTransactionFriend struct {
	Email      string
	Password   string
	FriendName string
}

type DeleteSplitTransactionReq struct {
	Email         string
	Password      string
	TransactionId uint
}

type DeleteSplitTransactionStringReq struct {
	Email       string
	Password    string
	TransString string
}

type GetSplitTransactionsReq struct {
	Email    string
	Password string
}

type GetUnSettledSplitTransactionsReq struct {
	Email    string
	Password string
}

type GetSettledSplitTransactionsReq struct {
	Email    string
	Password string
}

type GetFriendsReq struct {
	Email    string
	Password string
}

type TransactionAndSplitWithOneReq struct {
	Email       string
	Password    string
	Amount      int
	Category    string
	Description string
	SplitTag    string
	FriendName  string
	SplitAmount int
}

type MoneyLentFriend struct {
	Email      string
	Password   string
	FriendName string
}

type MoneyFriends struct {
	Email    string
	Password string
}

type FriendsMap struct {
	Email    string
	Password string
}
