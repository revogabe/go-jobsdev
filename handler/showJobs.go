package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShowJobsHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	println("===>", id)
	objectId, err := primitive.ObjectIDFromHex(id)

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "routeParameter").Error())
		return
	}
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	// Buscar o documento pelo ID
	filter := bson.M{"_id": objectId}
	var jobs schemas.Jobs
	db.Collection("jobs").FindOne(context.Background(), filter).Decode(&jobs)
	if err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "show-job", jobs)
}
