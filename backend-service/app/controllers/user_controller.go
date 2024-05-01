package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/faizauthar12/eccomerce/backend-service/app/models"
	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/backend-service/app/usecases"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Insert(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserController struct {
	ctx              context.Context
	mongod           mongodb.IMongoDB
	userUseCase      usecases.IUserUseCase
	userRepositories repositories.IUserRepository
}

func NewUserController(
	ctx context.Context,
	mongod mongodb.IMongoDB,
	userUseCase usecases.IUserUseCase,
) IUserController {
	return &UserController{
		ctx:         ctx,
		mongod:      mongod,
		userUseCase: userUseCase,
	}
}

func (c *UserController) Insert(ctx *gin.Context) {
	var result model.Response
	var user models.UserRequest

	err := ctx.BindJSON(&user)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	if len(user.Password) > 100 {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, "Password is too long")
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userData, errorLog := c.userUseCase.Insert(&user)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	userResponse, errorLog := c.userUseCase.GenerateToken(userData, os.Getenv("JWT_API_SECRET"))

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = userResponse
	result.StatusCode = http.StatusCreated

	ctx.JSON(http.StatusCreated, result)
}

func (c *UserController) Login(ctx *gin.Context) {
	var result model.Response
	var user models.UserLoginRequest

	err := ctx.BindJSON(&user)
	if err != nil {
		errorLog := helper.WriteLog(err, http.StatusBadRequest, err.Error())
		result.StatusCode = http.StatusBadRequest
		result.Error = errorLog
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	userData, errorLog := c.userUseCase.Authenticate(&user)

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	userResponse, errorLog := c.userUseCase.GenerateToken(userData, os.Getenv("JWT_API_SECRET"))

	if errorLog != nil {
		result.StatusCode = errorLog.StatusCode
		result.Error = errorLog
		ctx.JSON(errorLog.StatusCode, result)
		return
	}

	result.Data = userResponse
	result.StatusCode = http.StatusOK

	ctx.JSON(http.StatusOK, result)
}
