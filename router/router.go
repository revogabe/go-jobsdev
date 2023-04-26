package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"PUT", "DELETE"},
    AllowHeaders:     []string{"*"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
      return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
  }))

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