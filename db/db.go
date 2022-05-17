package db

import "github.com/fahrettinEmre/paymentSys/models"

var Account map[int]*models.Account = make(map[int]*models.Account)
var TransactionInfo map[int][]models.TransactionInfo = make(map[int][]models.TransactionInfo)
