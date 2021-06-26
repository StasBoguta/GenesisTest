package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getBtcRate(w http.ResponseWriter, r *http.Request){
	cookies, _ := r.Cookie("token")
	if cookies == nil || !validateJwt(cookies.Value){
		responseGenerator(w, "Unauthorized", 401)
		return
	}

	var bitcoin Bitcoin
	resp, _ := http.Get("https://api.kuna.io/v3/exchange-rates/btc")
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&bitcoin)
	uah := fmt.Sprintf("%.0f", bitcoin.Uah)

	response := BitcoinResponse{Course: uah}
	js, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
