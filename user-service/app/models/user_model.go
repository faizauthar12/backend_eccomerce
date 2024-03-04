package models

import "github.com/faizauthar12/backend_eccomerce/global-utils/model"

type User struct {
	UUID         string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	PasswordHash string `json:"password_hash,omitempty" bson:"password_hash,omitempty"`
	PasswordSalt string `json:"password_salt,omitempty" bson:"password_salt,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
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
	Name     string `json:"name" bson:"name,omitempty" binding:"required"`
	Email    string `json:"email" bson:"email,omitempty" binding:"required"`
	Password string `json:"password" bson:"password,omitempty" binding:"required"`
}
