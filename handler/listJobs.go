package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

func ListJobsHandler(ctx *gin.Context) {
	jobs := []schemas.Jobs{}

	if err := db.Find(&jobs).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error getting jobs")
		return 
	}

	sendSuccess(ctx, "list-jobs", jobs)
}