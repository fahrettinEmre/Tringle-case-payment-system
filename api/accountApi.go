package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fahrettinEmre/paymentSys/db"
	"github.com/fahrettinEmre/paymentSys/models"
	"github.com/fahrettinEmre/paymentSys/service"
	"github.com/gorilla/mux"
)

func Api(router *mux.Router) {
	router.HandleFunc("/account", CreateAccount).Methods("POST")
	router.HandleFunc("/account/{accountNumber}", getAccountInfo).Methods("GET")
	router.HandleFunc("/payment", payment).Methods("POST")
	router.HandleFunc("/deposit", deposit).Methods("POST")
	router.HandleFunc("/withdrow", withdrow).Methods("POST")
	router.HandleFunc("/accounting/{accountNumber}", transactionHistory).Methods("GET")

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("post request:createAccount")

	var acc models.Account
	err := json.NewDecoder(r.Body).Decode(&acc)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	acc.AccountNumber = len(db.Account)

	db.Account[acc.AccountNumber] = &acc
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

func getAccountInfo(w http.ResponseWriter, r *http.Request) {
	accountId, err := strconv.Atoi(mux.Vars(r)["accountNumber"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, exist := db.Account[accountId]
	if !exist {
		log.Println("account doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

func payment(w http.ResponseWriter, r *http.Request) {
	log.Println("post request:payment")
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sender, exist := db.Account[payment.SenderAccount]
	if !exist {
		log.Println("sender account doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	receiver, exist := db.Account[payment.ReceiverAccount]
	if !exist {
		log.Println("receiver account doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if payment.Amount > sender.Balance {
		log.Println("err: sender amount doesnt enough")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("err: sender amount doesnt enough"))
		return
	}
	sender.Balance = sender.Balance - payment.Amount
	receiver.Balance = receiver.Balance + payment.Amount

	db.Account[sender.AccountNumber] = sender
	db.Account[receiver.AccountNumber] = receiver

	service.SaveTransactionWithType(models.Pay, payment.Amount, sender.AccountNumber)

	log.Println("payment finished successfully")
	w.WriteHeader(http.StatusOK)

}

func withdrow(w http.ResponseWriter, r *http.Request) {
	log.Println("post request:withdrow")
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sender, exist := db.Account[payment.SenderAccount]
	if !exist {
		log.Println("sender account doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sender.Balance = sender.Balance - payment.Amount

	db.Account[sender.AccountNumber] = sender

	service.SaveTransactionWithType(models.With, payment.Amount, sender.AccountNumber)

	log.Println("Withdrow finished successfully")
	w.WriteHeader(http.StatusOK)

}

func deposit(w http.ResponseWriter, r *http.Request) {
	log.Println("post request:deposit")
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sender, exist := db.Account[payment.SenderAccount]
	if !exist {
		log.Println("sender account doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sender.Balance = sender.Balance + payment.Amount

	db.Account[sender.AccountNumber] = sender

	service.SaveTransactionWithType(models.Dep, payment.Amount, sender.AccountNumber)

	log.Println("Withdrow finished successfully")
	w.WriteHeader(http.StatusOK)

}

func transactionHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("request:transactionHistory")
	accountId, err := strconv.Atoi(mux.Vars(r)["accountNumber"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	infos, exist := db.TransactionInfo[accountId]

	if !exist {
		log.Println("transaction history doesnt exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(infos)
}
