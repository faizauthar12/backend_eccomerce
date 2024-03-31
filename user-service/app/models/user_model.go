package models

import "github.com/faizauthar12/backend_eccomerce/global-utils/model"

type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Name              string `json:"name" bson:"name"`
	Email             string `json:"email" bson:"email"`
	Role              string `json:"role" bson:"role"`
	PasswordHash      string `json:"password_hash" bson:"password_hash"`
	PasswordSalt      string `json:"password_salt" bson:"password_salt"`
	Token             string `json:"token" bson:"token"`
	RefreshToken      string `json:"refresh_token" bson:"refresh_token"`
	RefreshTokenCount int    `json:"refresh_token_count" bson:"refresh_token_count"`
	CreatedAt         int64  `json:"created_at" bson:"created_at"`
	UpdatedAt         int64  `json:"updated_at" bson:"updated_at"`
}

type UserChan struct {
	User       *User           `json:"user,omitempty"`
	Total      int64           `json:"total,omitempty"`
	Error      error           `json:"error,omitempty"`
	ErrorLog   *model.ErrorLog `json:"error_log,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
	StatusCode int             `json:"status_code,omitempty"`
}

type UserRequest struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Role     string `json:"role" bson:"role" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type UserResponse struct {
	UUID         string `json:"uuid" bson:"uuid"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Role         string `json:"role" bson:"role"`
	Token        string `json:"token" bson:"token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    int64  `json:"created_at" bson:"created_at"`
	UpdatedAt    int64  `json:"updated_at" bson:"updated_at"`
}
