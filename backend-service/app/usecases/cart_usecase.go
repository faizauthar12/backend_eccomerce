package usecases

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
)

type ICartUseCase interface {
	Get(request *models.Cart) (*models.Cart, *model.ErrorLog)
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

	insertCartChan := make(chan *models.CartChan)
	go u.cartRepository.InsertByUserUUID(request, u.ctx, insertCartChan)
	insertCart := <-insertCartChan

	if insertCart.Error != nil {
		return nil, insertCart.ErrorLog
	}

	return insertCart.Cart, nil
}
