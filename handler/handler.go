package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "POST jobs!",
	})
}

func ShowJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "GET jobs!",
	})
}

func DeleteJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "DELETE jobs!",
	})
}

func UpdateJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "PUT jobs!",
	})
}

func ListJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "GET all jobs!",
	})
}