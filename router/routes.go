package router

import (
	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1") 
	{
		v1.GET("/jobs", handler.ShowJobsHandler)
		v1.POST("/create", handler.CreateJobsHandler)
		v1.DELETE("/delete", handler.DeleteJobsHandler)
		v1.PUT("/update", handler.UpdateJobsHandler)
		v1.GET("/jobsall", handler.ListJobsHandler)
	}
}