package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/faizauthar12/eccomerce/backend-service/app/repositories"
	"github.com/faizauthar12/eccomerce/backend-service/app/usecases"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(
	mongod mongodb.IMongoDB,
	ctx context.Context,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := repositories.NewUserRepository(mongod)
		userUseCase := usecases.NewUserUseCase(userRepository, mongod)

		authorizationHeaderValue := c.GetHeader("Authorization")
		if len(authorizationHeaderValue) == 0 {
			err := helper.NewError("Authorization Headers is Required")
			errLog := &model.ErrorLog{
				Err:           err,
				Message:       err.Error(),
				SystemMessage: err.Error(),
				StatusCode:    http.StatusBadRequest,
			}

			response := &model.Response{
				Error:      errLog,
				StatusCode: errLog.StatusCode,
			}

			c.AbortWithStatusJSON(errLog.StatusCode, response)
			return
		}

		result := model.Response{}
		validateTokenWithClient, err := userUseCase.ParseAccessToken(authorizationHeaderValue, os.Getenv("JWT_API_SECRET"))

		if err != nil {
			errorLogData := helper.WriteLog(err, http.StatusInternalServerError, err.Error())
			result.Error = errorLogData
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Set("user", validateTokenWithClient)
		c.Next()
	}
}
