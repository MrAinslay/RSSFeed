package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MrAinslay/RSSFeed/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, usr database.User) {
	quer := r.URL.Query().Get("limit")
	lmt := 0
	var err error
	if quer == "" {
		lmt = 10
	} else {
		lmt, err = strconv.Atoi(quer)
		if err != nil {
			log.Printf("Error getting limit from url queries: %v", err)
			respondWithErr(w, 500, "Couldn't get limit from queries")
			return
		}
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: usr.ID,
		Limit:  int32(lmt),
	})
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		respondWithErr(w, 500, "Couldn't get posts")
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
