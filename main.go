package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/create", createUser).Methods("POST")
	myRouter.HandleFunc("/user/login", loginUser).Methods("POST")
	myRouter.HandleFunc("/btcRate", getBtcRate).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("API working")
	handleRequests()
}
