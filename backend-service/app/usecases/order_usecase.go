package usecases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrderUseCase interface {
	Get(request *models.OrderRequest) (*models.Orders, *model.ErrorLog)
	GetLastOrder(request *models.OrderRequest) (*models.Order, *model.ErrorLog)
	GetDetailOrder(uuid string) (*models.Order, *model.ErrorLog)
	Insert(request *models.Order) (*models.Order, *model.ErrorLog)
}

type OrderUseCase struct {
	orderRepository repositories.IOrderRepository
	mongod          mongodb.IMongoDB
	ctx             context.Context
}

func NewOrderUseCase(
	orderRepository repositories.IOrderRepository,
	mongod mongodb.IMongoDB,
	ctx context.Context,
) IOrderUseCase {
	return &OrderUseCase{
		orderRepository: orderRepository,
		mongod:          mongod,
		ctx:             ctx,
	}
}

func (u *OrderUseCase) Get(
	request *models.OrderRequest,
) (*models.Orders, *model.ErrorLog) {

	getOrderChan := make(chan *models.OrdersChan)
	go u.orderRepository.Get(request, u.ctx, getOrderChan)
	getOrderResult := <-getOrderChan

	if getOrderResult.Error != nil {
		return &models.Orders{}, getOrderResult.ErrorLog
	}

	if len(getOrderResult.Orders) == 0 {
		errorLogData := helper.WriteLog(mongo.ErrNoDocuments, http.StatusNotFound, nil)
		return &models.Orders{}, errorLogData
	}

	orders := &models.Orders{
		Orders: getOrderResult.Orders,
		Total:  getOrderResult.Total,
	}

	return orders, nil
}

func (u *OrderUseCase) GetDetailOrder(
	uuid string,
) (*models.Order, *model.ErrorLog) {

	getOrderChan := make(chan *models.OrderChan)
	go u.orderRepository.GetDetail(uuid, u.ctx, getOrderChan)
	getOrderResult := <-getOrderChan

	if getOrderResult.Error != nil {
		return &models.Order{}, getOrderResult.ErrorLog
	}

	return getOrderResult.Order, nil
}

func (u *OrderUseCase) GetLastOrder(
	request *models.OrderRequest,
) (*models.Order, *model.ErrorLog) {

	getOrderChan := make(chan *models.OrdersChan)
	go u.orderRepository.Get(request, u.ctx, getOrderChan)
	getOrderResult := <-getOrderChan

	fmt.Println("errorLog", getOrderResult.ErrorLog)

	if getOrderResult.Error != nil {
		return &models.Order{}, getOrderResult.ErrorLog
	}

	if len(getOrderResult.Orders) == 0 {
		errorLogData := helper.WriteLog(mongo.ErrNoDocuments, http.StatusNotFound, nil)
		return &models.Order{}, errorLogData
	}

	orders := &models.Order{
		UUID:      getOrderResult.Orders[0].UUID,
		UserUUID:  getOrderResult.Orders[0].UserUUID,
		OrderItem: getOrderResult.Orders[0].OrderItem,
		Total:     getOrderResult.Orders[0].Total,
		Address:   getOrderResult.Orders[0].Address,
		CreatedAt: getOrderResult.Orders[0].CreatedAt,
		UpdatedAt: getOrderResult.Orders[0].UpdatedAt,
	}

	return orders, nil
}

func (u *OrderUseCase) Insert(
	request *models.Order,
) (*models.Order, *model.ErrorLog) {

	insertOrderChan := make(chan *models.OrderChan)
	go u.orderRepository.Insert(request, u.ctx, insertOrderChan)
	insertOrderResult := <-insertOrderChan

	if insertOrderResult.Error != nil {
		return &models.Order{}, insertOrderResult.ErrorLog
	}

	return insertOrderResult.Order, nil
}
