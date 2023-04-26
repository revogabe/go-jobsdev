package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

func CreateJobsHandler(ctx *gin.Context) {
	request := CreateJobsRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("Error binding request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	jobs := schemas.Jobs{
		Title:       request.Title,
		Description: request.Description,
		Role:        request.Role,
		Company:     request.Company,
		Location:    request.Location,
		Remote:      request.Remote,
		Link:        request.Link,
		Experience:  request.Experience,
		Salary:      request.Salary,
		Approved:    false,
	}

	if _, err := db.Collection("jobs").InsertOne(context.Background(), jobs); err != nil {
		logger.Errorf("Error creating jobs: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating jobs on database")
		return
	}

	sendSuccess(ctx, "create-jobs", jobs)
}
