package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "GET all jobs!",
	})
}