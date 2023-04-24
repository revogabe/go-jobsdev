package config

import (
	"os"

	"github.com/revogabe/go-jobsdev/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database does not exist, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Jobs{})
	if err != nil {
		logger.Errorf("Failed to migrate database: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}