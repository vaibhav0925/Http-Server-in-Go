package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code> 499 {
		log.Printf("Log With 5XX error %v", msg)
	}
	type errRespose struct{
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errRespose{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "Application-json")
	w.WriteHeader(500)
	w.Write(dat)
}