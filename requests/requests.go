package requests

import "time"

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
	Time        time.Time
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
