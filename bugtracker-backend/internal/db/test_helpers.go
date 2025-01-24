package db

import (
	"os"
	"testing"

	"go.etcd.io/bbolt"
)

func SetupTestDB(t *testing.T) func() {
	// Use a temporary file for tests
	tmpFile := "test.db"
	os.Remove(tmpFile)
	
	// Override the database path for tests
	databasePath = tmpFile
	
	if err := Init(); err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Initialize counter bucket
	err := db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(counterBucket)
		if err != nil {
			return err
		}
		return b.Put([]byte("bug_id"), itob(0))
	})
	if err != nil {
		t.Fatalf("Failed to initialize counter: %v", err)
	}

	// Return cleanup function
	return func() {
		Cleanup()
		os.Remove(tmpFile)
	}
} 