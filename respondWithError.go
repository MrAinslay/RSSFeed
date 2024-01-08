package main

import (
	"log"
	"net/http"
)

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	type err struct {
		Error string `json:"error"`
	}

	rsp := err{
		Error: msg,
	}
	log.Printf("Responding with 5XX error: %s", msg)
	respondWithJSON(w, code, rsp)
}
