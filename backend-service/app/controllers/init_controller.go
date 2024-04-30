package controllers

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/backend-service/app/usecases"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
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

func InitHTTPProductController(
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IProductController {
	productRepository := repositories.NewProductRepository(mongod)
	productUseCase := usecases.NewProductUseCase(productRepository, mongod, ctx)
	handler := NewProductController(ctx, mongod, productUseCase, productRepository)
	return handler
}

func InitHTTPCartController(
	mongod mongodb.IMongoDB,
	ctx context.Context,
) ICartController {
	cartRepository := repositories.NewCartRepository(mongod)
	cartUseCase := usecases.NewCartUseCase(cartRepository, mongod, ctx)
	handler := NewCartController(ctx, mongod, cartUseCase, cartRepository)
	return handler
}
