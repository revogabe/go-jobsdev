package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/revogabe/go-jobsdev/schemas"
)

func AuthMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &schemas.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(*schemas.JwtClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		isAuthorized := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("username", claims.UserName)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func SignToken(user *schemas.TokenResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &schemas.JwtClaims{
		UserName: user.Name,
		Role:     user.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("Failed to sign token")
	}

	return tokenString, nil
}
