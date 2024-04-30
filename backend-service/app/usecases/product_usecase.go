package usecases

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
)

type IProductUseCase interface {
	Get(request *models.ProductRequest) (*models.Products, *model.ErrorLog)
	GetDetail(uuid string) (*models.Product, *model.ErrorLog)
	Insert(request *models.Product) (*models.Product, *model.ErrorLog)
	Update(request *models.Product) *model.ErrorLog
	Delete(request *models.Product) *model.ErrorLog
}

type ProductUseCase struct {
	productRepository repositories.IProductRepository
	mongod            mongodb.IMongoDB
	ctx               context.Context
}

func NewProductUseCase(
	productRepository repositories.IProductRepository,
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
		mongod:            mongod,
		ctx:               ctx,
	}
}

func (u *ProductUseCase) Get(
	request *models.ProductRequest,
) (*models.Products, *model.ErrorLog) {

	getProductChan := make(chan *models.ProductsChan)
	go u.productRepository.Get(request, u.ctx, getProductChan)
	getProductResult := <-getProductChan

	if getProductResult.Error != nil {
		return &models.Products{}, getProductResult.ErrorLog
	}

	products := &models.Products{
		Products: getProductResult.Products,
		Total:    getProductResult.Total,
	}

	return products, nil
}

func (u *ProductUseCase) GetDetail(
	uuid string,
) (*models.Product, *model.ErrorLog) {

	getDetailProductChan := make(chan *models.ProductChan)
	go u.productRepository.GetDetail(uuid, u.ctx, getDetailProductChan)
	getDetailProductResult := <-getDetailProductChan

	if getDetailProductResult.Error != nil {
		return &models.Product{}, getDetailProductResult.ErrorLog
	}

	return getDetailProductResult.Product, nil
}

func (u *ProductUseCase) Insert(
	request *models.Product,
) (*models.Product, *model.ErrorLog) {

	insertProductChan := make(chan *models.ProductChan)
	go u.productRepository.Insert(request, u.ctx, insertProductChan)
	insertProductResult := <-insertProductChan

	if insertProductResult.Error != nil {
		return &models.Product{}, insertProductResult.ErrorLog
	}

	return insertProductResult.Product, nil
}

func (u *ProductUseCase) Update(
	request *models.Product,
) *model.ErrorLog {

	updateProductChan := make(chan *models.ProductChan)
	go u.productRepository.Update(request, u.ctx, updateProductChan)
	updateProductResult := <-updateProductChan

	if updateProductResult.Error != nil {
		return updateProductResult.ErrorLog
	}

	return nil
}

func (u *ProductUseCase) Delete(
	request *models.Product,
) *model.ErrorLog {

	deleteProductChan := make(chan *models.ProductChan)
	go u.productRepository.Delete(request, u.ctx, deleteProductChan)
	deleteProductResult := <-deleteProductChan

	if deleteProductResult.Error != nil {
		return deleteProductResult.ErrorLog
	}

	return nil
}
