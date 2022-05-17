package models

type Transactions string

const (
	Pay  Transactions = "payment"
	Dep               = "deposit"
	With              = "withdrow"
)

type Transaction struct {
	AccounNumber int
	Amount       float32
	PaymentType  Transactions
}
