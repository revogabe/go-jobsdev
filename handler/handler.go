package handler

import (
	"github.com/revogabe/go-jobsdev/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger *config.Logger
	db     *mongo.Database
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMongoDB()
}
