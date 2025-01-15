package testutil

import (
	"os"
	"path/filepath"
)

const TestDBName = "test.db"

func GetTestDBPath() string {
	// Use temp directory for test database
	return filepath.Join(os.TempDir(), TestDBName)
}

func CleanupTestDB() error {
	return os.Remove(GetTestDBPath())
}
