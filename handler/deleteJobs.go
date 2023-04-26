package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

func DeleteJobsHandler(ctx *gin.Context) {
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

	if err := db.Find(&jobs, id).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("Error finding jobs with id %s on database", id))
		return
	}

	if err := db.Delete(&jobs).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting jobs with id %s on database", id))
		return
	}

	sendSuccess(ctx, "delete-jobs", jobs)
}