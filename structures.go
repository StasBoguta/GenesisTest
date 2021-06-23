package main

import "github.com/dgrijalva/jwt-go"

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

type UserStatusResponse struct{
	Message string `json:"message"`
}

type Users struct {
	Users []User `json:"users"`
}

type Bitcoin struct {
	Currency string `json:"currency"`
	Usd float32 `json:"usd"`
	Uah float32 `json:"uah"`
	Btc float32 `json:"btc"`
	Eur float32	`json:"eur"`
	Rub float32 `json:"rub"`
}

type BitcoinResponse struct {
	Course string `json:"course"`
}
