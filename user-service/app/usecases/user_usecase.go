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
	//Authenticate(request *models.UserRequest) (string, *model.ErrorLog)
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

	return inserUserResult.User, inserUserResult.ErrorLog
}

//func (u *UserUseCase) Authenticate(
//	request *models.UserRequest,
//) (string, *model.ErrorLog) {
//
//	authUserChan := make(chan *models.UserChan)
//	go u.userRepository.FindByEmail(request.Email, u.ctx, authUserChan)
//	authUserResult := <-authUserChan
//
//	if authUserResult.Error != nil {
//		return "", authUserResult.ErrorLog
//	}
//
//	insertedPasswordHash := hex.EncodeToString(
//		pbkdf2.Key(
//			[]byte(request.Password),
//			[]byte(authUserResult.User.PasswordSalt), 10000, 64, sha1.New),
//	)
//
//	if insertedPasswordHash != authUserResult.User.PasswordHash {
//		return "", &model.ErrorLog{
//			StatusCode: 401,
//			Message:    "Unauthorized",
//		}
//	}
//
//	return authUserResult.User, authUserResult.ErrorLog
//}
