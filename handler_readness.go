package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
	//responseWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}