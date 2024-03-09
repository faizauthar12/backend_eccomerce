package repositories

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/faizauthar12/backend_eccomerce/global-utils/helper"
	"github.com/faizauthar12/backend_eccomerce/global-utils/model"

	"github.com/dchest/uniuri"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/constants"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/models"
	log "github.com/sirupsen/logrus"
	"github.com/xdg-go/pbkdf2"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Insert(request *models.UserRequest, ctx context.Context, result chan *models.UserChan)
	FindByEmail(email string, ctx context.Context, result chan *models.UserChan)
}

type UserRepository struct {
	mongod         mongodb.IMongoDB
	logger         log.Logger
	collectionUser string
}

func NewUserRepository(mongod mongodb.IMongoDB) IUserRepository {
	return &UserRepository{
		mongod:         mongod,
		collectionUser: constants.COLLECTION_USER,
	}
}

func (r *UserRepository) Insert(
	request *models.UserRequest,
	ctx context.Context,
	result chan *models.UserChan,

) {
	timeNow := time.Now().Unix()

	response := &models.UserChan{}
	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	salt := uniuri.New()
	passwordHash := pbkdf2.Key([]byte(request.Password), []byte(salt), 10000, 64, sha1.New)

	user := models.User{
		UUID:         uniuri.New(),
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: hex.EncodeToString(passwordHash),
		PasswordSalt: salt,
		CreatedAt:    timeNow,
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		if mongoErr, ok := err.(mongo.WriteException); ok {
			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 { // 11000 is the code for a duplicate key error
					if strings.Contains(writeErr.Message, "email") {
						response.Error = errors.New(constants.ERROR_EMAIL_TAKEN)
					}
				}
			}
		}

		errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	response.User = &user
	result <- response
	return
}

func (r *UserRepository) FindByEmail(email string, ctx context.Context, result chan *models.UserChan) {
	response := &models.UserChan{}
	collection := r.mongod.Client().Database(constants.DATABASE).Collection(r.collectionUser)

	var user models.User
	err := collection.FindOne(ctx, models.User{Email: email}).Decode(&user)
	if err != nil {
		var errorLogData *model.ErrorLog
		if err == mongo.ErrNoDocuments {
			errorLogData = helper.WriteLog(err, http.StatusNotFound, err.Error())

		} else {
			errorLogData = helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		}

		response.Error = err
		response.ErrorLog = errorLogData
		result <- response
		return
	}

	response.User = &user
	result <- response
	return
}
