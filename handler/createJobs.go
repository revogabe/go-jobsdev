package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Criar uma instância vazia da estrutura Jobs
	jobs := schemas.Jobs{}

	// Definir valores dos campos da estrutura Jobs com base no objeto request
	jobs.Title = request.Title
	jobs.Description = request.Description
	jobs.Role = request.Role
	jobs.Company = request.Company
	jobs.Location = request.Location
	jobs.Remote = request.Remote
	jobs.Link = request.Link
	jobs.Experience = request.Experience
	jobs.Salary = request.Salary
	jobs.Approved = false

	// Inserir jobs no banco de dados
	res, err := db.Collection("jobs").InsertOne(context.Background(), jobs)
	if err != nil {
		logger.Errorf("Error creating jobs: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating jobs on database")
		return
	}

	// Extrair valor do atributo _id da resposta da função InsertOne
	id := res.InsertedID.(primitive.ObjectID)

	// Definir valor do campo _id da estrutura Jobs com o valor do atributo _id
	jobs.ID = id

	sendSuccess(ctx, "create-jobs", jobs)
}
