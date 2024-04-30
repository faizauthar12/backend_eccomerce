package controllers

import (
	"context"
	"net/http"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/backend-service/app/usecases"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

type ICartController interface {
	Get(ctx *gin.Context)
	Insert(ctx *gin.Context)
}

type CartController struct {
	ctx            context.Context
	mongod         mongodb.IMongoDB
	cartUseCase    usecases.ICartUseCase
	cartRepository repositories.ICartRepository
}

func NewCartController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	cartUseCase usecases.ICartUseCase,
	cartRepository repositories.ICartRepository,
) ICartController {
	return &CartController{
		ctx:            ctx,
		mongod:         mongod,
		cartUseCase:    cartUseCase,
		cartRepository: cartRepository,
	}
}

func (c *CartController) Get(ctx *gin.Context) {
	var result model.Response

	// var err error
	// var validationErrors []error

	cartRequest := &models.Cart{
		UserUUID: "WE NEED THE USER UUID FROM BEARER TOKEn",
	}

	cartResponse, errorLog := c.cartUseCase.Get(cartRequest)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = cartResponse
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)

}

func (c *CartController) Insert(ctx *gin.Context) {
	var result model.Response
	var cart models.Cart

	err := ctx.BindJSON(&cart)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	cartResponse, errorLog := c.cartUseCase.Insert(&cart)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = cartResponse
	result.StatusCode = http.StatusCreated

	ctx.JSON(http.StatusCreated, result)
}
