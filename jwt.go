package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type MyClaims struct {
	Login string
	jwt.StandardClaims
}

func generateJwt(login string) (string, time.Time){
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := MyClaims{
		login,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	return tokenString, expirationTime
}

func validateJwt(tokenString string) bool{
	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	_, ok := token.Claims.(*MyClaims)
	return ok
}
