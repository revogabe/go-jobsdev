package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	
  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  router.Use(cors.New(config))

	// Initialize Routes
	initializeRoutes(router)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Run the server
	router.Run("0.0.0.0:" + port)
}