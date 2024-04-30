package repositories

import (
	"context"
	"log"
	"net/http"

	"github.com/faizauthar12/backend_eccomerce/cart-service/app/constants"
	"github.com/faizauthar12/backend_eccomerce/cart-service/app/models"
	"github.com/faizauthar12/backend_eccomerce/global-utils/helper"
	"github.com/faizauthar12/backend_eccomerce/global-utils/model"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICartRepository interface {
	GetCartByUserUUID(userUUID string, ctx context.Context, result chan *models.CartChan)
	InsertByUserUUID(cart *models.Cart, ctx context.Context, result chan *models.CartChan)
}

type CartRepository struct {
	mongod            mongodb.IMongoDB
	logger            log.Logger
	collectionProduct string
}

func NewCartRepository(
	mongod mongodb.IMongoDB,
	logger log.Logger,
) ICartRepository {
	return &CartRepository{
		mongod:            mongod,
		logger:            logger,
		collectionProduct: constants.COLLECTION_CARTS,
	}
}

func (r *CartRepository) GetCartByUserUUID(
	userUUID string,
	ctx context.Context,
	result chan *models.CartChan,
) {
	response := &models.CartChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filterList := bson.D{bson.E{Key: "user_uuid", Value: userUUID}}

	var cart *models.Cart
	err := collection.FindOne(ctx, filterList).Decode(&cart)

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

	response.Cart = cart
	result <- response
	return
}

// Insert Will perform update with upsert option
func (r *CartRepository) InsertByUserUUID(
	cart *models.Cart,
	ctx context.Context,
	result chan *models.CartChan,
) {
	response := &models.CartChan{}

	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionProduct)

	filterList := bson.D{bson.E{Key: "user_uuid", Value: cart.UserUUID}}

	update := bson.D{{Key: "$set", Value: cart}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filterList, update, opts)
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
