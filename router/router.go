package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	router.Run("0.0.0.0:" + port)
}
