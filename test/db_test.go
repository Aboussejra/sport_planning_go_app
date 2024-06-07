package test

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize the test database using SQL file
func initTestDB(sqlFilePath string) (*gorm.DB, error) {
	// Open the in-memory SQLite database
	test_db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Read the SQL file
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return nil, err
	}

	// Execute the SQL commands
	if err := test_db.Exec(string(sqlBytes)).Error; err != nil {
		return nil, err
	}

	return test_db, nil
}

// Setup the test database
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := initTestDB("../database/init.sql")
	if err != nil {
		t.Fatalf("failed to initialize test database: %v", err)
	}
	return db
}
