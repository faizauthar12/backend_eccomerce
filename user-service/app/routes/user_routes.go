package routes

import (
	"context"

	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/controllers"
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
		userControllerGroup.POST("/user", ctrl.Insert)
	}

	// userControllerProtectedGroup := userControllerGroup.Use()
}
