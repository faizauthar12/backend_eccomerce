package repositories

import (
	"context"
	"net/http"
	"time"

	"github.com/faizauthar12/eccomerce/backend-service/app/constants"
	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOrderRepository interface {
	Get(request *models.OrderRequest, ctx context.Context, result chan *models.OrdersChan)
	GetDetail(uuid string, ctx context.Context, result chan *models.OrderChan)
	Insert(order *models.Order, ctx context.Context, result chan *models.OrderChan)
}

type OrderRepository struct {
	mongod            mongodb.IMongoDB
	collectionProduct string
}

func NewOrderRepository(
	mongod mongodb.IMongoDB,
) IOrderRepository {
	return &OrderRepository{
		mongod:            mongod,
		collectionProduct: constants.COLLECTION_ORDERS,
	}
}

func (r *OrderRepository) Get(
	request *models.OrderRequest,
	ctx context.Context,
	result chan *models.OrdersChan,
) {

	response := &models.OrdersChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	// Make filterList for query and assign it to be empty bson.D instead of nil
	filterList := bson.D{}

	filterList = append(filterList, bson.E{Key: "user_uuid", Value: request.UserUUID})

	skipOptions := options.Find().SetLimit(request.NumItems).SetSkip((request.Pages - 1) * request.NumItems)

	orderBy := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	mongoCursor, err := collection.Find(
		ctx,
		filterList,
		skipOptions,
		orderBy,
	)

	if err != nil {
		var errorLogData *model.ErrorLog

		if err == mongo.ErrNoDocuments {
			errorLogData = helper.WriteLog(err, http.StatusNotFound, nil)
		} else {
			errorLogData = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		}

		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	defer mongoCursor.Close(ctx)

	orders := []*models.Order{}
	for mongoCursor.Next(ctx) {
		order := &models.Order{}
		err := mongoCursor.Decode(order)
		if err != nil {
			errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			response.Error = err
			response.ErrorLog = errorLogData
			result <- response
			return
		}

		orders = append(orders, order)
	}

	response.Orders = orders
	result <- response
	return
}

func (r *OrderRepository) GetDetail(
	uuid string,
	ctx context.Context,
	result chan *models.OrderChan,
) {

	response := &models.OrderChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filterList := bson.D{{Key: "uuid", Value: uuid}}

	mongoResult := collection.FindOne(ctx, filterList)

	if mongoResult.Err() != nil {
		var errorLogData *model.ErrorLog

		if mongoResult.Err() == mongo.ErrNoDocuments {
			errorLogData = helper.WriteLog(mongoResult.Err(), http.StatusNotFound, nil)
		} else {
			errorLogData = helper.WriteLog(mongoResult.Err(), http.StatusInternalServerError, mongoResult.Err().Error())
		}

		response.Error = mongoResult.Err()
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	order := &models.Order{}
	err := mongoResult.Decode(order)
	if err != nil {
		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	response.Order = order
	result <- response
	return
}

func (r *OrderRepository) Insert(
	order *models.Order,
	ctx context.Context,
	result chan *models.OrderChan,
) {
	response := &models.OrderChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	var total float64
	for indexItem, item := range order.OrderItem {
		order.OrderItem[indexItem].Total = item.Price * float64(item.Quantity)
		total += order.OrderItem[indexItem].Total
	}

	// insertedOrder := models.Order{
	// 	UUID:      uuid.New().String(),
	// 	UserUUID:  order.UserUUID,
	// 	OrderItem: order.OrderItem,
	// 	Total:     total,
	// 	CreatedAt: time.Now().Unix(),
	// 	Address:   order.Address,
	// }

	order.UUID = uuid.New().String()
	order.Total = total
	order.CreatedAt = time.Now().Unix()

	_, err := collection.InsertOne(ctx, &order)
	if err != nil {
		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	response.Order = order
	result <- response
	return
}
