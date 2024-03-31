package usecases

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/faizauthar12/backend_eccomerce/global-utils/helper"
	"github.com/faizauthar12/backend_eccomerce/global-utils/model"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/constants"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/models"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xdg-go/pbkdf2"
)

type IUserUseCase interface {
	Insert(request *models.UserRequest) (*models.User, *model.ErrorLog)
	Authenticate(request *models.UserLoginRequest) (*models.User, *model.ErrorLog)
	GenerateToken(request *models.User, apiSecret string) (*models.UserResponse, *model.ErrorLog)
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

func (u *UserUseCase) Authenticate(
	request *models.UserLoginRequest,
) (*models.User, *model.ErrorLog) {

	authUserChan := make(chan *models.UserChan)
	go u.userRepository.FindByEmail(request.Email, u.ctx, authUserChan)
	authUserResult := <-authUserChan

	if authUserResult.Error != nil {
		return &models.User{}, authUserResult.ErrorLog
	}

	insertedPasswordHash := hex.EncodeToString(
		pbkdf2.Key(
			[]byte(request.Password),
			[]byte(authUserResult.User.PasswordSalt), 10000, 64, sha1.New),
	)

	if insertedPasswordHash != authUserResult.User.PasswordHash {
		return &models.User{}, &model.ErrorLog{
			StatusCode: 401,
			Message:    "Unauthorized",
		}
	}

	return authUserResult.User, authUserResult.ErrorLog
}

func (u *UserUseCase) GenerateToken(
	request *models.User,
	apiSecret string,
) (*models.UserResponse, *model.ErrorLog) {

	timeNow := time.Now()
	timeNowUnix := timeNow.Unix()

	newRefreshTokenCount := request.RefreshTokenCount + 1

	if request.RefreshTokenCount == 4 {
		refreshClaims := jwt.RegisteredClaims{
			Issuer: constants.JWT_TOKEN_ISSUEER,
			IssuedAt: &jwt.NumericDate{
				Time: timeNow,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: timeNow.Add(constants.JWT_REFRESH_TOKEN_LIFESPAN * time.Hour),
			},
		}

		// Generate Refresh Token
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
		refreshTokenString, err := refreshToken.SignedString([]byte(apiSecret))

		if err != nil {
			errorLog := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			return nil, errorLog
		}

		newRefreshTokenCount = 0
		request.RefreshToken = refreshTokenString
	}

	userClaims := &models.UserJWT{
		UUID:              request.UUID,
		Name:              request.Name,
		Role:              request.Role,
		Email:             request.Email,
		RefreshTokenCount: newRefreshTokenCount,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: constants.JWT_TOKEN_ISSUEER,
			IssuedAt: &jwt.NumericDate{
				Time: timeNow,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: timeNow.Add(constants.JWT_TOKEN_LIFESPAN * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(apiSecret))
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
		return nil, errorLog
	}

	// assign the new refresh token count
	request.RefreshTokenCount = newRefreshTokenCount

	// update the unix timestamp
	request.UpdatedAt = timeNowUnix

	// update the mongodb user document
	updateUserChan := make(chan *models.UserChan)
	go u.userRepository.Update(request, u.ctx, updateUserChan)
	updateUserResult := <-updateUserChan

	if updateUserResult.Error != nil {
		return nil, updateUserResult.ErrorLog
	}

	userResponse := &models.UserResponse{
		UUID:         request.UUID,
		Name:         request.Name,
		Role:         request.Role,
		Email:        request.Email,
		Token:        tokenString,
		RefreshToken: request.RefreshToken,
		CreatedAt:    request.CreatedAt,
		UpdatedAt:    request.UpdatedAt,
	}

	return userResponse, nil
}
