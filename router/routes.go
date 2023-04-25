package router

import (
	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath) 
	{
		v1.GET("/jobs", handler.ShowJobsHandler)
		v1.POST("/create", handler.CreateJobsHandler)
		v1.PUT("/approved", handler.ApprovedJobsHandler)
		v1.DELETE("/delete", handler.DeleteJobsHandler)
		v1.GET("/jobsall", handler.ListJobsHandler)
	}
}