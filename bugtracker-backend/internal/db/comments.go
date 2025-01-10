package db

import (
	"bugtracker-backend/internal/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

func CreateComment(bugID string, comment *models.Comment) error {
	// Convert string ID to int
	idInt, err := strconv.Atoi(bugID)
	if err != nil {
		return fmt.Errorf("invalid bug ID format")
	}

	// Check if bug exists
	_, err = GetBug(idInt)
	if err != nil {
		return fmt.Errorf("bug not found")
	}

	comment.ID = uuid.New().String()
	comment.CreatedAt = time.Now()
	comment.BugID = bugID

	return database.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(commentsBucket)
		encoded, err := json.Marshal(comment)
		if err != nil {
			return err
		}
		return b.Put([]byte(comment.ID), encoded)
	})
}

func GetComments(bugID string) ([]models.Comment, error) {
	var comments []models.Comment

	err := database.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(commentsBucket)
		return b.ForEach(func(k, v []byte) error {
			var comment models.Comment
			if err := json.Unmarshal(v, &comment); err != nil {
				return err
			}
			if comment.BugID == bugID {
				comments = append(comments, comment)
			}
			return nil
		})
	})

	return comments, err
}
