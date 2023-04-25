package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revogabe/go-jobsdev/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"errorCode": code,
	})
}


func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string `json:"message"`
	Data schemas.JobsResponse `json:"data"`	
}


