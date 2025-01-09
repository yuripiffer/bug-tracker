package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
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
}

func CreateBugHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBugRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bug := &models.Bug{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.CreateBug(bug); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	bug, err := db.GetBug(id)
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

	var req models.CreateBugRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve existing bug
	bug, err := db.GetBug(id)
	if err != nil {
		http.Error(w, "Bug not found", http.StatusNotFound)
		return
	}

	// Update fields
	bug.Title = req.Title
	bug.Description = req.Description
	bug.Status = req.Status
	bug.Priority = req.Priority

	// Save updated bug
	if err := db.CreateBug(bug); err != nil { // Reuse CreateBug for simplicity
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bug)
}

// DeleteBugHandler deletes a bug by its ID
func DeleteBugHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := database.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)
		return b.Delete([]byte(id))
	})

	if err != nil {
		http.Error(w, "Failed to delete bug", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}