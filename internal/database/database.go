package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Wikid82/CaddyProxyManagerPlus/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize sets up the database connection and runs migrations
func Initialize(dataPath string) error {
	// Create data directory if it doesn't exist
	if err := os.MkdirAll(dataPath, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	dbPath := filepath.Join(dataPath, "caddyproxymanager.db")
	
	// Open database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	// Run migrations
	if err := runMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

// runMigrations automatically migrates the schema
func runMigrations() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.ProxyHost{},
		&models.Settings{},
		&models.AccessLog{},
		&models.CrowdSecDecision{},
	)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
