package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MrAinslay/RSSFeed/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, usr database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		log.Printf("Couldnt decode request: %v", err)
		respondWithErr(w, 500, "Couldn't decode request")
		return
	}

	timeNow := time.Now()
	u, err := uuid.NewV7()
	if err != nil {
		log.Printf("Error creating uuid: %v", err)
		respondWithErr(w, 500, "Couldn't create new uuid")
		return
	}

	fd, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        u,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    usr.ID,
		Url:       params.Url,
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("Error creating new feed: %v", err)
		respondWithErr(w, 500, "Couldn't create new feed")
		return
	}

	newU, err := uuid.NewV7()
	if err != nil {
		log.Printf("Error creating uuid: %v", err)
		respondWithErr(w, 500, "Couldn't create new uuid")
		return
	}

	cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        newU,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    usr.ID,
		FeedID:    u,
	})

	respondWithJSON(w, 200, databaseFeedToFeed(fd))
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	dbFds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		log.Printf("Error while getting feeds: %v", err)
		respondWithErr(w, 500, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(dbFds))
}
