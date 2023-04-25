package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateJobsRequest true "Request body"
// @Success 200 {object} CreateJobsRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /create [post]

func CreateJobsHandler(ctx *gin.Context) {
	request := CreateJobsRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Erro validate jobs: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	jobs := schemas.Jobs{
		Title: request.Title,
		Description: request.Description,
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: request.Remote,
		Link: request.Link,
		Experience: request.Experience,
		Salary: request.Salary,
	}

	if err := db.Create(&jobs).Error; err != nil {
		logger.Errorf("Erro create jobs: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating jobs on database")
		return
	}

	sendSuccess(ctx, "create-jobs", jobs)
}