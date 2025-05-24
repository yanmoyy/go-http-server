package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yanmoyy/go-http-server/internal/api"
	"github.com/yanmoyy/go-http-server/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env: %s", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("Platform must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET "+api.HealthzEndpoint, handlerReadiness)
	mux.HandleFunc("POST "+api.UsersEndpoint, apiCfg.handleCreateUser)
	mux.HandleFunc("POST "+api.ChirpsEndpoint, apiCfg.handleCreateChirp)
	mux.HandleFunc("GET "+api.ChirpsEndpoint, apiCfg.handleGetChirpList)
	mux.HandleFunc("GET "+api.ChirpsEndpoint+"/{"+api.ChripIDParam+"}", apiCfg.handleGetChirpByID)
	mux.HandleFunc("GET "+api.MetricsEndpoint, apiCfg.handlerMetrics)
	mux.HandleFunc("POST "+api.ResetEndpoint, apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
