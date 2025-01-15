package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"bugtracker-backend/internal/db"
	"bugtracker-backend/internal/models"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/bugs", CreateBug).Methods("POST")
	r.HandleFunc("/api/bugs", GetBugs).Methods("GET")
	r.HandleFunc("/api/bugs/{id}", GetBug).Methods("GET")
	r.HandleFunc("/api/bugs/{id}", UpdateBug).Methods("PUT")
	r.HandleFunc("/api/bugs/{id}", DeleteBug).Methods("DELETE")
	RegisterCommentRoutes(r)
	r.HandleFunc("/api/health", HealthCheck).Methods("GET")
}

func CreateBug(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBugRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode create bug request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "title is required",
		})
		return
	}

	bug := &models.Bug{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.CreateBug(bug); err != nil {
		log.Printf("Failed to create bug: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bug)
}

func GetBugs(w http.ResponseWriter, r *http.Request) {
	bugs, err := db.GetAllBugs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bugs)
}

func GetBug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid bug ID",
		})
		return
	}

	bug, err := db.GetBug(idInt)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "bug not found" {
			status = http.StatusNotFound
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(bug)
}

func UpdateBug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Convert string ID to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req models.CreateBugRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve existing bug
	existingBug, err := db.GetBug(idInt)
	if err != nil {
		http.Error(w, "Bug not found", http.StatusNotFound)
		return
	}

	// Update fields
	existingBug.Title = req.Title
	existingBug.Description = req.Description
	existingBug.Status = req.Status
	existingBug.Priority = req.Priority

	// Use UpdateBug instead of CreateBug
	if err := db.UpdateBug(existingBug); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingBug)
}

func DeleteBug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	log.Printf("Attempting to delete bug with ID: %d", idInt)

	if err := db.DeleteBug(idInt); err != nil {
		log.Printf("Failed to delete bug with ID %d: %v", idInt, err)
		http.Error(w, "Failed to delete bug", http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully deleted bug with ID: %d", idInt)
	w.WriteHeader(http.StatusNoContent)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
