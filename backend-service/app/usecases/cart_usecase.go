package usecases

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
)

type ICartUseCase interface {
	Get(request *models.Cart) (*models.Cart, *model.ErrorLog)
	GetTotalItemInCart(request *models.Cart) (int, *model.ErrorLog)
	Insert(request *models.Cart) (*models.Cart, *model.ErrorLog)
}

type CartUseCase struct {
	cartRepository repositories.ICartRepository
	mongod         mongodb.IMongoDB
	ctx            context.Context
}

func NewCartUseCase(
	cartRepository repositories.ICartRepository,
	mongod mongodb.IMongoDB,
	ctx context.Context,
) ICartUseCase {
	return &CartUseCase{
		cartRepository: cartRepository,
		mongod:         mongod,
		ctx:            ctx,
	}
}

func (u *CartUseCase) Get(
	request *models.Cart,
) (*models.Cart, *model.ErrorLog) {

	getCartChan := make(chan *models.CartChan)
	go u.cartRepository.GetCartByUserUUID(request.UserUUID, u.ctx, getCartChan)
	getCart := <-getCartChan

	if getCart.Error != nil {
		return nil, getCart.ErrorLog
	}

	return getCart.Cart, nil
}

func (u *CartUseCase) Insert(
	request *models.Cart,
) (*models.Cart, *model.ErrorLog) {

	getCartChan := make(chan *models.CartChan)
	go u.cartRepository.GetCartByUserUUID(request.UserUUID, u.ctx, getCartChan)
	getCart := <-getCartChan

	if getCart.ErrorLog != nil {
		fmt.Println("error", getCart.ErrorLog)
		if getCart.ErrorLog.StatusCode == http.StatusInternalServerError {
			return nil, getCart.ErrorLog
		}
	}

	if getCart.Cart == nil {
		request.CreatedAt = time.Now().Unix()
	} else {
		request.CreatedAt = getCart.Cart.CreatedAt
		request.UpdatedAt = time.Now().Unix()
	}

	insertCartChan := make(chan *models.CartChan)
	go u.cartRepository.InsertByUserUUID(request, u.ctx, insertCartChan)
	insertCart := <-insertCartChan

	if insertCart.Error != nil {
		return nil, insertCart.ErrorLog
	}

	return insertCart.Cart, nil
}

func (u *CartUseCase) GetTotalItemInCart(request *models.Cart) (int, *model.ErrorLog) {

	getCartChan := make(chan *models.CartChan)
	go u.cartRepository.GetCartByUserUUID(request.UserUUID, u.ctx, getCartChan)
	getCart := <-getCartChan

	if getCart.Error != nil {
		return 0, getCart.ErrorLog
	}

	return len(getCart.Cart.CartItems), nil
}
