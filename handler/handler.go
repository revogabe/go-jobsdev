package handler

import (
	"github.com/revogabe/go-jobsdev/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db 	 	*gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySQL()
}







