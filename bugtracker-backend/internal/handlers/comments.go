package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bugtracker-backend/internal/db"
	"bugtracker-backend/internal/models"
)

func RegisterCommentRoutes(r *mux.Router) {
	r.HandleFunc("/api/bugs/{bugId}/comments", CreateComment).Methods("POST")
	r.HandleFunc("/api/bugs/{bugId}/comments", GetComments).Methods("GET")
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bugID := vars["bugId"]

	var req models.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if req.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "content is required",
		})
		return
	}

	comment := &models.Comment{
		Author:  req.Author,
		Content: req.Content,
	}

	if err := db.CreateComment(bugID, comment); err != nil {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bugID := vars["bugId"]

	w.Header().Set("Content-Type", "application/json")

	// Validate bug ID format
	if _, err := strconv.Atoi(bugID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid bug ID",
		})
		return
	}

	comments, err := db.GetComments(bugID)
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

	json.NewEncoder(w).Encode(comments)
}
