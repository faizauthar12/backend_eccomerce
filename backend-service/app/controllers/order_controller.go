package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/usecases"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

type IOrderController interface {
	Get(ctx *gin.Context)
	GetLastOrder(ctx *gin.Context)
	GetDetailOrder(ctx *gin.Context)
	Insert(ctx *gin.Context)
}

type OrderController struct {
	ctx          context.Context
	mongod       mongodb.IMongoDB
	orderUseCase usecases.IOrderUseCase
}

func NewOrderController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	orderUseCase usecases.IOrderUseCase,
) IOrderController {
	return &OrderController{
		ctx:          ctx,
		mongod:       mongod,
		orderUseCase: orderUseCase,
	}
}

func (c *OrderController) Get(ctx *gin.Context) {
	result := model.Response{}

	numItemsString, numItemsStringExist := ctx.GetQuery("num_items")
	if !numItemsStringExist {
		numItemsString = "10"
	}

	pagesString, pagesStringExist := ctx.GetQuery("pages")
	if !pagesStringExist {
		pagesString = "1"
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

	user := ctx.Value("user").(*models.UserJWT)

	orderRequest := &models.OrderRequest{
		NumItems: numItems,
		Pages:    pages,
		UserUUID: user.UUID,
	}

	orderResponse, errorLog := c.orderUseCase.Get(orderRequest)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orderResponse
	result.Pages = pages
	result.NumItems = numItems
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}

func (c *OrderController) GetLastOrder(ctx *gin.Context) {
	result := model.Response{}

	user := ctx.Value("user").(*models.UserJWT)

	orderRequest := &models.OrderRequest{
		UserUUID: user.UUID,
	}

	orderResponse, errorLog := c.orderUseCase.GetLastOrder(orderRequest)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orderResponse
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}

func (c *OrderController) GetDetailOrder(ctx *gin.Context) {
	result := model.Response{}

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

	orderResponse, errorLog := c.orderUseCase.GetDetailOrder(uuid)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orderResponse
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}

func (c *OrderController) Insert(ctx *gin.Context) {
	result := model.Response{}
	order := models.Order{}

	err := ctx.BindJSON(&order)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	user := ctx.Value("user").(*models.UserJWT)
	order.UserUUID = user.UUID

	orderResponse, errorLog := c.orderUseCase.Insert(&order)
	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = orderResponse
	result.StatusCode = http.StatusCreated

	ctx.JSON(http.StatusCreated, result)
}
