package main

import (
	"net/http"
)

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 400, "Something went wrong")
	//responseWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}