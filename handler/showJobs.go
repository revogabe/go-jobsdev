package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ShowJobsHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "routeParameter").Error())
		return
	}

	// Obter o cliente do MongoDB a partir do contexto
	client, ok := ctx.Value("mongo").(*mongo.Client)
	if !ok {
		sendError(ctx, http.StatusInternalServerError, "Error getting MongoDB client from context")
		return
	}

	// Obter a coleção de jobs
	collection := client.Database("jobsdb").Collection("jobs")

	// Buscar o documento pelo ID
	filter := bson.M{"_id": id}
	var jobs schemas.Jobs
	err := collection.FindOne(context.Background(), filter).Decode(&jobs)
	if err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	sendSuccess(ctx, "show-job", jobs)
}
