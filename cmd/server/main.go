package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

type Response struct {
	Message string `json:"message"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func main() {
	log.Printf("User service is listening on: %v\n", PORT)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := Response{
			Message: "Hello from User Service API!",
			Service: "calorie-snap-user-svc",
			Version: "v1.0.0",
		}

		json.NewEncoder(w).Encode(response)
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", PORT),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
