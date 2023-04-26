package config

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	config := cors.DefaultConfig()
  config.AllowOrigins = []string{"*"}

	// Initialize SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("failed to initialize SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}