package usecases

import (
	"context"

	"github.com/faizauthar12/backend_eccomerce/global-utils/model"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/models"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/repositories"
)

type IUserUseCase interface {
	Insert(request *models.UserRequest) (*models.User, *model.ErrorLog)
}

type UserUseCase struct {
	userRepository repositories.IUserRepository
	mongod         mongodb.IMongoDB
	ctx            context.Context
}

func NewUserUseCase(
	userRepository repositories.IUserRepository,
	mongod mongodb.IMongoDB,
) IUserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		mongod:         mongod,
	}
}

func (u *UserUseCase) Insert(
	request *models.UserRequest,
) (*models.User, *model.ErrorLog) {

	inserUserChan := make(chan *models.UserChan)
	go u.userRepository.Insert(request, u.ctx, inserUserChan)
	inserUserResult := <-inserUserChan

	if inserUserResult.Error != nil {
		return &models.User{}, inserUserResult.ErrorLog
	}

	return inserUserResult.User, &model.ErrorLog{}
}
