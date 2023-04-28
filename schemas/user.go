package schemas

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

type TokenResponse struct {
	Name string `bson:"username"`
	Role string `bson:"role"`
}

type UserResponse struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Role  string             `json:"role" bson:"role"`
	Token string             `json:"token" bson:"token"`
}
