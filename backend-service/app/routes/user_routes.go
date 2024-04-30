package routes

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/controllers"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongod mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPUserController(mongod, ctx)

	userControllerGroup := g.Group(path)
	{
		userControllerGroup.POST("", ctrl.Insert)
		userControllerGroup.POST("/login", ctrl.Login)
	}

	// userControllerProtectedGroup := userControllerGroup.Use()
}
