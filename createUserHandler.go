package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MrAinslay/RSSFeed/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) crtUsrHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("Error decoding request %v", err)
		respondWithErr(w, 500, "Couldn't decode request")
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	u, _ := uuid.NewV7()
	timeNow := time.Now()

	usr, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        u,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("Error creating user: %v", err)
		respondWithErr(w, 500, "Couldn't create user")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(usr))
}

func (cfg *apiConfig) getUserByKey(w http.ResponseWritter r *http.Request) {
 
}
