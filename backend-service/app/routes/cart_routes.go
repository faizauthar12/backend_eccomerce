package routes

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/controllers"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func InitCartRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongod mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPCartController(mongod, ctx)

	cartControllerGroup := g.Group(path)
	{
		cartControllerGroup.POST("", ctrl.Insert)
		cartControllerGroup.GET("", ctrl.Get)
	}
}
