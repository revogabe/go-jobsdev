package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func DeleteJobsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "DELETE jobs!",
	})
}