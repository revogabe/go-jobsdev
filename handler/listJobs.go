package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func ListJobsHandler(ctx *gin.Context) {
	jobs := []schemas.Jobs{}

	cur, err := db.Collection("jobs").Find(ctx, bson.M{})
	if err != nil {
		logger.Errorf("Error getting jobs: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error getting jobs from database")
		return
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var job schemas.Jobs
		if err := cur.Decode(&job); err != nil {
			logger.Errorf("Error decoding job: %v", err.Error())
			sendError(ctx, http.StatusInternalServerError, "Error getting jobs from database")
			return
		}

		jobs = append(jobs, job)
	}

	if err := cur.Err(); err != nil {
		logger.Errorf("Error getting jobs: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error getting jobs from database")
		return
	}

	sendSuccess(ctx, "list-jobs", jobs)
}
