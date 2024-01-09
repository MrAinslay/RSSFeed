package main

import (
	"log"
	"net/http"

	"github.com/MrAinslay/RSSFeed/internal/auth"
	"github.com/MrAinslay/RSSFeed/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			log.Printf("Error finding valid api key: %v", err)
			respondWithErr(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		usr, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			log.Printf("Error finding user: %v", err)
			respondWithErr(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		handler(w, r, usr)
	}
}
