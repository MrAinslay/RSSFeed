package main

import "net/http"

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 500, "Internal Server Error")
}
