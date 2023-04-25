package router

import (
	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/handler"

	docs "github.com/revogabe/go-jobsdev/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath) 
	{
		v1.GET("/jobs", handler.ShowJobsHandler)
		v1.POST("/create", handler.CreateJobsHandler)
		v1.DELETE("/delete", handler.DeleteJobsHandler)
		v1.GET("/jobsall", handler.ListJobsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}