package handler

import (
	"github.com/gin-gonic/gin"
)

type UserValidationResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func UserCheckHandler(ctx *gin.Context) {
	username := ctx.GetString("username")
	role := ctx.GetString("role")

	response := UserValidationResponse{
		Username: username,
		Role:     role,
	}

	sendSuccess(ctx, "show-job", response)
	return
}
