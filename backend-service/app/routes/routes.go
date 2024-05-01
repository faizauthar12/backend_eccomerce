package routes

import (
	"context"
	"net/http"

	"github.com/faizauthar12/eccomerce/backend-service/app/middlewares"
	"github.com/faizauthar12/eccomerce/global-utils/model"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func InitHTTPRoute(
	g *gin.Engine,
	mongodbClient mongodb.IMongoDB,
	ctx context.Context,
) {

	g.Use(middlewares.CORSMiddleware(), middlewares.JSONMiddleware())

	g.GET("/health-check", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{"status": "OK"})
	})

	InitUserRoute("/user", ctx, g, mongodbClient)

	InitProductRoute("/product", ctx, g, mongodbClient)

	InitCartRoute("/cart", ctx, g, mongodbClient)

	InitOrderRoute("/order", ctx, g, mongodbClient)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Error: &model.ErrorLog{
				Message:       "Not Found",
				SystemMessage: "Not Found",
			},
		})
	})
}
