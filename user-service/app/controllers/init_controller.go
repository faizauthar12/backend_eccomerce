package controllers

import (
	"context"

	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/repositories"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/usecases"
)

func InitHTTPUserController(
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IUserController {
	userRepository := repositories.NewUserRepository(mongod)
	userUseCase := usecases.NewUserUseCase(userRepository, mongod)
	handler := NewUserController(ctx, mongod, userUseCase)
	return handler
}
