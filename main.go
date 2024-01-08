package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/MrAinslay/RSSFeed/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load("keys.env")

	db, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Handle("/", http.FileServer(http.Dir(".")))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", errHandler)

	v1Router.Post("/users", apiCfg.crtUsrHandler)
	v1Router.Get("/users", apiCfg.getUserByKey)

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
