package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "GET jobs!",
	})
}