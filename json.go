package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	rsp := payload

	dat, err := json.Marshal(&rsp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(status)
	w.Write(dat)
}

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
