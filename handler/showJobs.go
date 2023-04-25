package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

func ShowJobsHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	jobs := schemas.Jobs{}

	if err := db.First(&jobs, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "show-job", jobs)
}