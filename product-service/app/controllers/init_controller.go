package controllers

import (
	"context"

	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/repositories"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/usecases"
)

func InitHTTPProductController(
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IProductController {
	productRepository := repositories.NewProductRepository(mongod)
	productUseCase := usecases.NewProductUseCase(productRepository, mongod, ctx)
	handler := NewProductController(ctx, mongod, productUseCase, productRepository)
	return handler
}
