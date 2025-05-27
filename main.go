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
	jwtSecret      string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env: %s", err)
	}

	dbURL := os.Getenv(envDBURL)
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv(envPlatform)
	if platform == "" {
		log.Fatal("Platform must be set")
	}
	jwtSecret := os.Getenv(envJWTSecret)
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env is not set")
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
		jwtSecret:      jwtSecret,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))

	mux.HandleFunc("GET "+api.EndpointHealthz, handlerReadiness)

	mux.HandleFunc("POST "+api.EndpointUsers, apiCfg.handleCreateUser)
	mux.HandleFunc("PUT "+api.EndpointUsers, apiCfg.handleUpdateUser)
	mux.HandleFunc("POST "+api.EndpointLogin, apiCfg.handleLogin)

	mux.HandleFunc("POST "+api.EndpointChirps, apiCfg.handleCreateChirp)
	mux.HandleFunc("GET "+api.EndpointChirps, apiCfg.handleGetChirpList)
	mux.HandleFunc("GET "+api.EndpointChirps+"/{"+api.ChirpIDParam+"}", apiCfg.handleGetChirpByID)
	mux.HandleFunc("DELETE "+api.EndpointChirps+"/{"+api.ChirpIDParam+"}", apiCfg.handleDeleteChirpByID)
	mux.HandleFunc("GET "+api.EndpointMetrics, apiCfg.handlerMetrics)
	mux.HandleFunc("POST "+api.EndpointReset, apiCfg.handlerReset)

	// Token
	mux.HandleFunc("POST "+api.EndpointRefresh, apiCfg.handleRefresh)
	mux.HandleFunc("POST "+api.EndpointRevoke, apiCfg.handleRevoke)

	// Polka Webhooks
	mux.HandleFunc("POST "+api.EndpointPolkaWebhooks, apiCfg.handlePolkaWebhooks)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
