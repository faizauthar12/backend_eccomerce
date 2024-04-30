package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/faizauthar12/eccomerce/backend-service/app/routes"
	"github.com/faizauthar12/eccomerce/global-utils/middlewares"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/gin-gonic/gin"
)

func MainHTTPHandler(
	mongodbClient mongodb.IMongoDB,
	ctx context.Context,
) {

	g := gin.Default()
	g.Use(middlewares.CORSMiddleware(), middlewares.JSONMiddleware(), middlewares.RequestIdMiddleware())

	routes.InitHTTPRoute(g, mongodbClient, ctx)

	addr := fmt.Sprintf(":%s", os.Getenv("MAIN_PORT"))

	g.Run(addr)
}
