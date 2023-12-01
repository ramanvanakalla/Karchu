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

type NetAmountByCategory struct {
	Email    string
	Password string
}
