package main

import (
	"fmt"
	"net/http"

	"github.com/fahrettinEmre/paymentSys/api"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("test")
	router := mux.NewRouter().StrictSlash(true)
	api.Api(router)

	http.ListenAndServe(":5050", router)

}
