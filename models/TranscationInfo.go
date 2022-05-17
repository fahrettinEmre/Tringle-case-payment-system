package models

type TransactionInfo struct {
	AccountNumber   int
	Amount          float32
	TransactionType Transactions
	CreateAt        int
}
