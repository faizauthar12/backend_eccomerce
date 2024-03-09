package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/faizauthar12/backend_eccomerce/global-utils/helper"
	"github.com/faizauthar12/backend_eccomerce/global-utils/model"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/models"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/repositories"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/usecases"
	"github.com/gin-gonic/gin"
)

type IProductController interface {
	Get(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ProductController struct {
	ctx                 context.Context
	mongod              mongodb.IMongoDB
	productUseCase      usecases.IProductUseCase
	productRepositories repositories.IProductRepository
}

func NewProductController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	productUseCase usecases.IProductUseCase,
	productRepositories repositories.IProductRepository,
) IProductController {
	return &ProductController{
		ctx:                 ctx,
		mongod:              mongod,
		productUseCase:      productUseCase,
		productRepositories: productRepositories,
	}
}

func (c *ProductController) Get(ctx *gin.Context) {
	var result model.Response

	numItemsString, numItemsStringExist := ctx.GetQuery("num_items")
	if !numItemsStringExist {
		numItemsString = "10"
	}

	pagesString, pagesStringExist := ctx.GetQuery("pages")
	if !pagesStringExist {
		pagesString = "1"
	}

	nameString, nameStringExist := ctx.GetQuery("name")
	if !nameStringExist {
		nameString = ""
	}

	priceGteString, priceGteStringExist := ctx.GetQuery("price_gte")
	if !priceGteStringExist {
		priceGteString = "0"
	}

	priceLteString, priceLteStringExist := ctx.GetQuery("price_lte")
	if !priceLteStringExist {
		priceLteString = "0"
	}

	numItems, err := strconv.ParseInt(numItemsString, 10, 64)
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		errorLog := helper.WriteLog(err, http.StatusBadRequest, "Parameter 'num_items' harus bernilai integer")
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	pages, err := strconv.ParseInt(pagesString, 10, 64)
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		errorLog := helper.WriteLog(err, http.StatusBadRequest, "Parameter 'pages' harus bernilai integer")
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	priceGte, err := strconv.ParseInt(priceGteString, 10, 64)
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		errorLog := helper.WriteLog(err, http.StatusBadRequest, "Parameter 'price_gte' harus bernilai integer")
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	priceLte, err := strconv.ParseInt(priceLteString, 10, 64)
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		errorLog := helper.WriteLog(err, http.StatusBadRequest, "Parameter 'price_lte' harus bernilai integer")
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	productRequest := models.ProductRequest{
		NumItems: numItems,
		Pages:    pages,
		Name:     nameString,
		PriceGte: priceGte,
		PriceLte: priceLte,
	}

	productsResponse, errorLog := c.productUseCase.Get(&productRequest)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = productsResponse
	result.Pages = pages
	result.NumItems = numItems
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)

}

func (c *ProductController) GetDetail(ctx *gin.Context) {
	var result model.Response

	uuid := ctx.Param("uuid")
	var err error
	var validationErrors []error

	if uuid == "" {
		err = helper.NewError("UUID is required")
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		errorMessages := []string{}
		for _, v := range validationErrors {
			errorMessages = append(errorMessages, v.Error())
		}
		err := helper.NewError("Validation Error")
		errorLog := helper.WriteLog(err, http.StatusUnprocessableEntity, errorMessages)
		result.Error = errorLog
		result.StatusCode = http.StatusUnprocessableEntity
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	productResponse, errorLog := c.productUseCase.GetDetail(uuid)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = productResponse
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}

func (c *ProductController) Insert(ctx *gin.Context) {
	var result model.Response
	var product models.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	productResponse, errorLog := c.productUseCase.Insert(&product)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = productResponse
	result.StatusCode = http.StatusCreated

	ctx.JSON(http.StatusCreated, result)
}

func (c *ProductController) Update(ctx *gin.Context) {
	var result model.Response
	var product models.Product

	uuid := ctx.Param("uuid")
	var err error
	var validationErrors []error

	if uuid == "" {
		err = helper.NewError("UUID is required")
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		errorMessages := []string{}
		for _, v := range validationErrors {
			errorMessages = append(errorMessages, v.Error())
		}
		err := helper.NewError("Validation Error")
		errorLog := helper.WriteLog(err, http.StatusUnprocessableEntity, errorMessages)
		result.Error = errorLog
		result.StatusCode = http.StatusUnprocessableEntity
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	err = ctx.BindJSON(&product)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	// Assign the UUID the variable
	product.UUID = uuid

	errorLog := c.productUseCase.Update(&product)
	if errorLog != nil {
		result.StatusCode = http.StatusInternalServerError
		result.Error = errorLog
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	var result model.Response
	var product models.Product

	uuid := ctx.Param("uuid")
	var err error
	var validationErrors []error

	if uuid == "" {
		err = helper.NewError("UUID is required")
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		errorMessages := []string{}
		for _, v := range validationErrors {
			errorMessages = append(errorMessages, v.Error())
		}
		err := helper.NewError("Validation Error")
		errorLog := helper.WriteLog(err, http.StatusUnprocessableEntity, errorMessages)
		result.Error = errorLog
		result.StatusCode = http.StatusUnprocessableEntity
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	// Assign the UUID to the variable
	product.UUID = uuid

	errorLog := c.productUseCase.Delete(&product)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}
