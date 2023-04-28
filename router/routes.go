package router

import (
	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/handler"
	authMiddleware "github.com/revogabe/go-jobsdev/middlewares/jwt"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	baseAuthPath := "/api/v1/auth"
	v1 := router.Group(basePath)
	authRoute := router.Group(baseAuthPath)

	authRoute.Use(authMiddleware.AuthMiddleware([]string{"admin", "owner"}))

	{
		v1.GET("/jobs", handler.ShowJobsHandler)
		v1.POST("/create", handler.CreateJobsHandler)
		v1.GET("/jobsall", handler.ListJobsHandler)
		v1.POST("/login", handler.LoginHandler)
		authRoute.PUT("/approved", handler.ApprovedJobsHandler)
		authRoute.DELETE("/delete", handler.DeleteJobsHandler)
		authRoute.GET("/validate", handler.UserCheckHandler)
	}
}
