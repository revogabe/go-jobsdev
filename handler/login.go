package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	authMiddleware "github.com/revogabe/go-jobsdev/middlewares/jwt"
	"github.com/revogabe/go-jobsdev/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("Error binding request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	user := schemas.User{}
	err := db.Collection("users").FindOne(context.Background(), bson.M{"username": request.Username}).Decode(&user)
	if err != nil {
		logger.Errorf("Error finding user: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logger.Errorf("Error comparing passwords: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	tokenRequest := schemas.TokenResponse{}

	tokenRequest.Name = user.Name
	tokenRequest.Role = user.Role

	tokenString, err := authMiddleware.SignToken(&tokenRequest)
	if err != nil {
		logger.Errorf("Error generating JWT: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error generating JWT")
		return
	}

	response := schemas.UserResponse{}

	response.ID = user.ID
	response.Token = tokenString
	response.Role = user.Role
	response.Name = user.Name

	sendSuccess(ctx, "login", response)
}
