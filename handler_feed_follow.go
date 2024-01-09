package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MrAinslay/RSSFeed/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerAddFeedFollow(w http.ResponseWriter, r *http.Request, usr database.User) {
	type rqstParameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	rqst := rqstParameters{}
	if err := decoder.Decode(&rqst); err != nil {
		log.Printf("Error decoding request: %v", err)
		respondWithErr(w, 500, "Couldn't decode request")
		return
	}

	timeNow := time.Now()
	u, err := uuid.NewV7()
	if err != nil {
		log.Printf("Error creating new uuid: %v", err)
		respondWithErr(w, 500, "Couldn't create new uuid")
		return
	}

	fdFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        u,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    usr.ID,
		FeedID:    rqst.FeedId,
	})
	if err != nil {
		log.Printf("Error creating new feed follow: %v", err)
		respondWithErr(w, 500, "Couldn't create new feed follow")
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowToFeedFollow(fdFollow))
}

func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, usr database.User) {
	id := chi.URLParam(r, "feedFollowID")
	uId, _ := uuid.Parse(id)
	_, err := cfg.DB.DeleteFeedFollow(r.Context(), uId)
	log.Println(uId, id)
	if err != nil {
		log.Printf("Error deleting feed follow: %v", err)
		respondWithErr(w, 500, "Couldn't delete feed follow")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, usr database.User) {
	fdFollows, err := cfg.DB.GetFeedFollows(r.Context(), usr.ID)
	if err != nil {
		log.Printf("Error getting feed follows: %v", err)
		respondWithErr(w, 500, "Couldn't get feed follows")
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(fdFollows))
}
