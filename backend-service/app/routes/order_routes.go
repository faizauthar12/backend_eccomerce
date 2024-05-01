package routes

import (
	"context"

	"github.com/faizauthar12/eccomerce/backend-service/app/controllers"
	"github.com/faizauthar12/eccomerce/backend-service/app/middlewares"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func InitOrderRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongod mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPOrderController(mongod, ctx)

	orderControllerGroup := g.Group(path)
	{
		orderControllerGroup.Use(middlewares.TokenMiddleware(mongod, ctx))
		orderControllerGroup.GET("", ctrl.Get)
		orderControllerGroup.GET("/last", ctrl.GetLastOrder)
		orderControllerGroup.GET(":uuid", ctrl.GetDetailOrder)
		orderControllerGroup.POST("", ctrl.Insert)
	}
}
