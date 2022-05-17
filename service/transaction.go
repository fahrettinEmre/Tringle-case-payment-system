package service

import (
	"time"

	"github.com/fahrettinEmre/paymentSys/db"
	"github.com/fahrettinEmre/paymentSys/models"
)

func SaveTransactionWithType(t models.Transactions, amount float32, accountNumber int) {
	transactionInfo := models.TransactionInfo{}
	transactionInfo.TransactionType = t
	transactionInfo.AccountNumber = accountNumber
	transactionInfo.Amount = amount
	transactionInfo.CreateAt = time.Now().Nanosecond()

	account := db.TransactionInfo[accountNumber]
	account = append(account, transactionInfo)
	db.TransactionInfo[accountNumber] = account

}
