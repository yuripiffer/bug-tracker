package db

import (
	"os"
	"testing"

	"bugtracker-backend/internal/models"
	"bugtracker-backend/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseInitialization(t *testing.T) {
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	err := Init()
	assert.NoError(t, err)
	defer func() {
		CleanupTestDB()
		Cleanup()
	}()

	// Test DB by creating a bug
	bug := &models.Bug{Title: "Test", Description: "Test"}
	err = CreateBug(bug)
	assert.NoError(t, err)
}

func TestMultipleInitializations(t *testing.T) {
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	// First initialization
	err := Init()
	assert.NoError(t, err)
	Cleanup()

	// Second initialization should work
	err = Init()
	assert.NoError(t, err)
	defer func() {
		CleanupTestDB()
		Cleanup()
	}()
}

func TestCleanup(t *testing.T) {
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	err := Init()
	assert.NoError(t, err)
	Cleanup()

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
	defer func() {
		CleanupTestDB()
		Cleanup()
	}()
}

func TestConcurrentInitializations(t *testing.T) {
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	err := Init()
	assert.NoError(t, err)
	defer func() {
		CleanupTestDB()
		Cleanup()
	}()

	// Attempt second initialization without cleanup
	err = Init()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database already initialized")
}

func TestMain(m *testing.M) {
	// Set test mode
	os.Setenv("TEST_MODE", "1")

	// Run tests
	code := m.Run()

	// Cleanup
	testutil.CleanupTestDB()

	os.Exit(code)
}
