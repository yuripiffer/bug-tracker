package db

import (
	"encoding/json"
	"fmt"
	"log"

	"bugtracker-backend/internal/models"
	"go.etcd.io/bbolt"
)

var (
	database    *bbolt.DB
	bugsBucket  = []byte("bugs")
	databasePath = "bugs.db" // Ensure this path is correct
)

// Init initializes the BoltDB database
func Init() error {
	// Open the BoltDB database file
	db, err := bbolt.Open(databasePath, 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	database = db

	// Create the bugs bucket if it doesn't exist
	err = database.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bugsBucket)
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	log.Println("Database initialized successfully.")
	return nil
}

// CreateBug inserts a new bug into the database
func CreateBug(bug *models.Bug) error {
	return database.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)

		// Serialize the bug to JSON
		encoded, err := json.Marshal(bug)
		if err != nil {
			return fmt.Errorf("failed to marshal bug: %w", err)
		}

		// Insert the bug with its ID as the key
		return b.Put([]byte(bug.ID), encoded)
	})
}

// GetBug retrieves a bug by its ID
func GetBug(id string) (*models.Bug, error) {
	var bug models.Bug

	err := database.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)
		data := b.Get([]byte(id))
		if data == nil {
			return fmt.Errorf("bug not found")
		}

		return json.Unmarshal(data, &bug)
	})

	if err != nil {
		return nil, err
	}

	return &bug, nil
}

// GetAllBugs retrieves all bugs from the database
func GetAllBugs() ([]*models.Bug, error) {
	var bugs []*models.Bug

	err := database.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)

		return b.ForEach(func(k, v []byte) error {
			var bug models.Bug
			if err := json.Unmarshal(v, &bug); err != nil {
				return fmt.Errorf("failed to unmarshal bug %s: %w", k, err)
			}
			bugs = append(bugs, &bug)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return bugs, nil
}

// Cleanup closes the database
func Cleanup() {
	if database != nil {
		err := database.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database closed successfully.")
		}
	}
}