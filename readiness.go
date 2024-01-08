package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	type readinessRsp struct {
		Status string `json:"status"`
	}

	rsp := readinessRsp{
		Status: "ok",
	}

	respondWithJSON(w, 200, rsp)
}
