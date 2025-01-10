package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"bugtracker-backend/internal/models"

	"go.etcd.io/bbolt"
)

var (
	database       *bbolt.DB
	bugsBucket     = []byte("bugs")
	commentsBucket = []byte("comments")
	counterBucket  = []byte("counter")
	databasePath   = "bugs.db" // Ensure this path is correct
)

// Init initializes the BoltDB database
func Init() error {
	// Open the BoltDB database file
	db, err := bbolt.Open(databasePath, 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	database = db

	// Create both buckets if they don't exist
	err = database.Update(func(tx *bbolt.Tx) error {
		// Create bugs bucket
		_, err := tx.CreateBucketIfNotExists(bugsBucket)
		if err != nil {
			return fmt.Errorf("create bugs bucket: %w", err)
		}

		// Create comments bucket
		_, err = tx.CreateBucketIfNotExists(commentsBucket)
		if err != nil {
			return fmt.Errorf("create comments bucket: %w", err)
		}

		// Create counter bucket
		_, err = tx.CreateBucketIfNotExists(counterBucket)
		if err != nil {
			return fmt.Errorf("create counter bucket: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create buckets: %w", err)
	}

	log.Println("Database initialized successfully.")
	return nil
}

// CreateBug inserts a new bug into the database
func CreateBug(bug *models.Bug) error {
	return database.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)

		// Get next ID
		nextID, err := getNextID(tx)
		if err != nil {
			return err
		}

		bug.ID = nextID

		// Serialize the bug to JSON
		encoded, err := json.Marshal(bug)
		if err != nil {
			return fmt.Errorf("failed to marshal bug: %w", err)
		}

		// Insert the bug with its ID as the key
		return b.Put(itob(bug.ID), encoded)
	})
}

// GetBug retrieves a bug by its ID
func GetBug(id int) (*models.Bug, error) {
	var bug models.Bug

	err := database.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)
		data := b.Get(itob(id))
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

// DeleteBug removes a bug from the database by its ID
func DeleteBug(id int) error {
	return database.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bugsBucket)
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.Delete(itob(id))
	})
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

// Add this function to get the next ID
func getNextID(tx *bbolt.Tx) (int, error) {
	b := tx.Bucket(counterBucket)
	id := b.Get([]byte("lastBugID"))

	var nextID int
	if id == nil {
		nextID = 1
	} else {
		nextID = btoi(id) + 1
	}

	if err := b.Put([]byte("lastBugID"), itob(nextID)); err != nil {
		return 0, err
	}

	return nextID, nil
}

// Helper functions for converting between int and []byte
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
