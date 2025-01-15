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
	r.HandleFunc("/api/bugs", CreateBugHandler).Methods("POST")
	r.HandleFunc("/api/bugs", GetBugsHandler).Methods("GET")
	r.HandleFunc("/api/bugs/{id}", GetBugHandler).Methods("GET")
	r.HandleFunc("/api/bugs/{id}", UpdateBugHandler).Methods("PUT")
	r.HandleFunc("/api/bugs/{id}", DeleteBugHandler).Methods("DELETE")
	RegisterCommentRoutes(r)
}

func CreateBugHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBugRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode create bug request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		log.Printf("Failed to create bug with ID %d: %v", bug.ID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully created bug with ID: %d", bug.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bug)
}

func GetBugsHandler(w http.ResponseWriter, r *http.Request) {
	bugs, err := db.GetAllBugs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bugs)
}

func GetBugHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	bug, err := db.GetBug(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bug)
}

// UpdateBugHandler updates an existing bug
func UpdateBugHandler(w http.ResponseWriter, r *http.Request) {
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

// DeleteBugHandler deletes a bug by its ID
func DeleteBugHandler(w http.ResponseWriter, r *http.Request) {
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
