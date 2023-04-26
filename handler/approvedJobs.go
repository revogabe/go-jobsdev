package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApprovedJobsHandler(ctx *gin.Context) {
	request := ApprovedJobsRequest{}

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

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"approved": true}}

	var jobs schemas.Jobs

	if err := db.Collection("jobs").FindOneAndUpdate(ctx, filter, update).Decode(&jobs); err != nil {
		if err == mongo.ErrNoDocuments {
			sendError(ctx, http.StatusNotFound, "Job not found")
			return
		}
		logger.Errorf("Error updating job: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating job")
		return
	}

	if !request.Approved {
		jobs.Approved = true
	} else {
		jobs.Approved = false
	}

	if _, err := db.Collection("jobs").ReplaceOne(ctx, filter, jobs); err != nil {
		logger.Errorf("Error replacing job: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating job")
		return
	}

	sendSuccess(ctx, "approved-jobs", jobs)
}
