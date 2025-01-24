package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("Health check request received from %s", r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Sending health check response")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
} 