package db

import (
	"os"
	"testing"

	"bugtracker-backend/internal/models"

	"github.com/stretchr/testify/assert"
)

const testDBPath = "test.db"

func init() {
	// Set test database path for all tests
	os.Setenv("DB_PATH", testDBPath)
}

func cleanup() {
	Cleanup()
	// Remove test database file after each test
	os.Remove(testDBPath)
}

func TestDatabaseInitialization(t *testing.T) {
	err := Init()
	assert.NoError(t, err)
	defer cleanup()

	// Test DB by creating a bug
	bug := &models.Bug{Title: "Test", Description: "Test"}
	err = CreateBug(bug)
	assert.NoError(t, err)
}

func TestMultipleInitializations(t *testing.T) {
	// First initialization
	err := Init()
	assert.NoError(t, err)

	// Cleanup after first initialization
	cleanup()

	// Second initialization should work
	err = Init()
	assert.NoError(t, err)
	defer cleanup()
}

func TestCleanup(t *testing.T) {
	err := Init()
	assert.NoError(t, err)
	cleanup()

	// Test DB is inaccessible after cleanup
	bug := &models.Bug{Title: "Test", Description: "Test"}
	err = CreateBug(bug)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database not initialized",
		"Should get 'database not initialized' error after cleanup")
}

func TestInitWithInvalidPath(t *testing.T) {
	// Save original values
	originalPath := os.Getenv("DB_PATH")
	originalDBPath := databasePath
	defer func() {
		os.Setenv("DB_PATH", originalPath)
		databasePath = originalDBPath
	}()

	// Set invalid path
	invalidPath := "/invalid/path/db.sqlite"
	t.Setenv("DB_PATH", invalidPath)
	databasePath = invalidPath

	err := Init()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open database")
	defer cleanup()
}

func TestConcurrentInitializations(t *testing.T) {
	// First initialization
	err := Init()
	assert.NoError(t, err)
	defer cleanup()

	// Attempt second initialization without cleanup
	err = Init()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database already initialized")
}
