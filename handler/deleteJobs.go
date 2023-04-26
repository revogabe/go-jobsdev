package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteJobsHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	authorization := ctx.GetHeader("authorization")
	secretKey := os.Getenv("SECRET_KEY")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	if authorization != secretKey {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("authorization", "GetHeader").Error())
		return
	}

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	filter := bson.M{"_id": objectId}
	var jobs schemas.Jobs

	err = db.Collection("jobs").FindOne(context.Background(), filter).Decode(&jobs)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sendError(ctx, http.StatusNotFound, fmt.Sprintf("Jobs with id %s not found", id))
			return
		}
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error finding jobs with id %s on database", id))
		return
	}

	_, err = db.Collection("jobs").DeleteOne(context.Background(), filter)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting jobs with id %s on database", id))
		return
	}

	sendSuccess(ctx, "delete-jobs", jobs)
}
