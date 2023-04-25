package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"

	"github.com/joho/godotenv"
)

func ApprovedJobsHandler(ctx *gin.Context) {
	request := ApprovedJobsRequest{}

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	id := ctx.Query("id")
	authorization := ctx.GetHeader("authorization")
	secretKey := os.Getenv("SECRET_KEY")

	if authorization != secretKey {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("authorization", "GetHeader").Error())
		return
	}

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	jobs := schemas.Jobs{}

	if err := db.First(&jobs, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	// Update opening
	if !request.Approved {
		jobs.Approved = true
	}

	// Save opening
	if err := db.Save(&jobs).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "approved-jobs", jobs)
}