package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"hive-honeypot-network/internal/database"
	"hive-honeypot-network/internal/events"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status:  "ok",
		Service: "HIVE API",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode health response: %v", err)
	}
}

func databaseHealthHandler(dbPing func(context.Context) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		if err := dbPing(ctx); err != nil {
			log.Printf("Database health check failed: %v", err)
			http.Error(w, "Database unavailable", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(map[string]string{
			"status":   "ok",
			"database": "connected",
		}); err != nil {
			log.Printf("Failed to encode database health response: %v", err)
		}
	}
}

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// Existing health endpoints
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/health/db", databaseHealthHandler(db.PingContext))

	// Event ingestion and retrieval endpoints
	store := database.NewDBStore(db)
	eventHandler := events.EventHandler{Store: store}

	mux.HandleFunc("/api/v1/events", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			eventHandler.List(w, r)
		case http.MethodPost:
			eventHandler.Create(w, r)
		default:
			w.Header().Set("Allow", "GET, POST")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"error": "Method not allowed",
			})
		}
	})

	addr := ":8080"
	log.Printf("HIVE API starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}

}
