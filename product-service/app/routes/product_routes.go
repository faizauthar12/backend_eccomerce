package routes

import (
	"context"

	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/product-service/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitProductRoute(
	path string,
	ctx context.Context,
	g *gin.Engine,
	mongod mongodb.IMongoDB,
) {
	ctrl := controllers.InitHTTPProductController(mongod, ctx)

	productControllerGroup := g.Group(path)
	{
		productControllerGroup.POST("", ctrl.Insert)
		productControllerGroup.GET("", ctrl.Get)
		productControllerGroup.GET(":uuid", ctrl.GetDetail)
		productControllerGroup.PATCH(":uuid", ctrl.Update)
		productControllerGroup.DELETE(":uuid", ctrl.Delete)
	}
}
