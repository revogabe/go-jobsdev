package config

import (
	"github.com/revogabe/go-jobsdev/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	dsn := "root:ovxujh2xyL5YefpDPTqi@tcp(containers-us-west-47.railway.app:7812)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	// Create DB and Connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Check if DB exists
	err = db.AutoMigrate(&schemas.Jobs{})
	if err != nil {
		logger.Errorf("Failed to migrate database: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
