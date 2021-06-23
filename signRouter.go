package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"regexp"
)


func createUser(w http.ResponseWriter, r *http.Request){
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u.Login == "" || u.Password == ""{
		responseGenerator(w, "Bad credentials", http.StatusBadRequest)
		return
	}

	re:=regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z]+\\.[a-zA-Z]+")
	if !re.MatchString(u.Login){
		responseGenerator(w, "Bad credentials", http.StatusBadRequest)
		return
	}

	u.Password = encodeData(u.Password)
	u.Login = encodeData(u.Login)

	w.Header().Set("Content-Type", "application/json")
	if addNewUser(u){
		responseGenerator(w, "User was been created", 200)
	} else{
		responseGenerator(w,"User with same login already exists", 409)
	}
}

func responseGenerator(w http.ResponseWriter, message string, statusCode int){
	response := UserStatusResponse{
		message,
	}
	js, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Write(js)
}

func encodeData(data string) string{
	return b64.StdEncoding.EncodeToString([]byte(data))
}

func loginUser(w http.ResponseWriter, r *http.Request){
	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	creds.Password = encodeData(creds.Password)
	creds.Login = encodeData(creds.Login)

	if !validateCredentials(creds){
		w.WriteHeader(401)
		return
	}

	tokenString, expirationTime := generateJws(creds.Login)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path: "/",
	})

	response := UserStatusResponse{
		"Logged in",
	}
	js, _ := json.Marshal(response)
	w.Write(js)
}
