package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bugtracker-backend/internal/db"
	"bugtracker-backend/internal/models"
)

func RegisterCommentRoutes(r *mux.Router) {
	r.HandleFunc("/api/bugs/{bugId}/comments", CreateCommentHandler).Methods("POST")
	r.HandleFunc("/api/bugs/{bugId}/comments", GetCommentsHandler).Methods("GET")
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bugID := vars["bugId"]

	var req models.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode comment request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment := &models.Comment{
		Author:  req.Author,
		Content: req.Content,
	}

	log.Printf("Creating comment for bug %s by %s", bugID, req.Author)

	if err := db.CreateComment(bugID, comment); err != nil {
		log.Printf("Failed to create comment: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully created comment with ID: %s", comment.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bugID := vars["bugId"]

	comments, err := db.GetComments(bugID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
