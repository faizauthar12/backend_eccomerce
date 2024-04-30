package repositories

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/faizauthar12/eccomerce/backend-service/app/constants"
	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
)

type IProductRepository interface {
	Get(request *models.ProductRequest, ctx context.Context, result chan *models.ProductsChan) // Get all products
	GetDetail(uuid string, ctx context.Context, result chan *models.ProductChan)               // Get product detail
	Insert(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Insert new product
	Update(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Update product
	Delete(product *models.Product, ctx context.Context, result chan *models.ProductChan)      // Delete product
}

type ProductRepository struct {
	mongod            mongodb.IMongoDB
	logger            log.Logger
	collectionProduct string
}

func NewProductRepository(
	mongod mongodb.IMongoDB,
) IProductRepository {
	return &ProductRepository{
		mongod:            mongod,
		collectionProduct: constants.COLLECTION_PRODUCT,
	}
}

func (r *ProductRepository) Get(
	request *models.ProductRequest,
	ctx context.Context,
	result chan *models.ProductsChan,
) {

	response := &models.ProductsChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	// Make filterList for query and assign it to be empty bson.D instead of nil
	filterList := bson.D{}

	if request.Name != "" {
		filterList = append(filterList, bson.E{Key: "name", Value: primitive.Regex{Pattern: request.Name, Options: "i"}})
	}

	if request.PriceGte > 0 {
		filterList = append(filterList, bson.E{Key: "price", Value: bson.D{{Key: "$gte", Value: request.PriceGte}}})
	}

	if request.PriceLte > 0 {
		filterList = append(filterList, bson.E{Key: "price", Value: bson.D{{Key: "$lte", Value: request.PriceLte}}})
	}

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

	var products []*models.Product
	for mongoCursor.Next(ctx) {
		var product models.Product
		err = mongoCursor.Decode(&product)
		if err != nil {
			errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			response.Error = err
			response.ErrorLog = errorLogData
			result <- response
			return
		}

		products = append(products, &product)
	}

	response.Products = products
	result <- response
	return
}

func (r *ProductRepository) GetDetail(
	uuid string,
	ctx context.Context,
	result chan *models.ProductChan,
) {

	response := &models.ProductChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filter := bson.D{{Key: "uuid", Value: uuid}}

	var product models.Product
	err := collection.FindOne(ctx, filter).Decode(&product)

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

	response.Product = &product
	result <- response
	return
}

func (r *ProductRepository) Insert(
	product *models.Product,
	ctx context.Context,
	result chan *models.ProductChan,
) {

	response := &models.ProductChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	insertedProduct := models.Product{
		UUID:      uuid.New().String(),
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: time.Now().Unix(),
	}

	_, err := collection.InsertOne(ctx, insertedProduct)
	if err != nil {

		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	response.Product = &insertedProduct
	result <- response
	return
}

func (r *ProductRepository) Update(
	product *models.Product,
	ctx context.Context,
	result chan *models.ProductChan,
) {

	response := &models.ProductChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filter := bson.D{{Key: "uuid", Value: product.UUID}}

	update := bson.D{{Key: "$set", Value: product}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	// response.Product = product
	result <- response
	return
}

func (r *ProductRepository) Delete(
	product *models.Product,
	ctx context.Context,
	result chan *models.ProductChan,
) {

	response := &models.ProductChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filter := bson.D{{Key: "uuid", Value: product.UUID}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	result <- response
	return
}
