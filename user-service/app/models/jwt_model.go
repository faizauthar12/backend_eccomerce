package models

import "github.com/golang-jwt/jwt/v5"

type UserJWT struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Name              string `json:"name" bson:"name"`
	Email             string `json:"email" bson:"email"`
	Role              string `json:"role" bson:"role"`
	RefreshTokenCount int    `json:"refreshtokencount"`
	jwt.RegisteredClaims
}
