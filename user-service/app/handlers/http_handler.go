package handlers

import (
	"context"
	"fmt"
	"github.com/faizauthar12/backend_eccomerce/global-utils/mongodb"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/middlewares"
	"github.com/faizauthar12/backend_eccomerce/user-service/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
)

func MainHttpHandler(
	mongodbClient mongodb.IMongoDB,
	ctx context.Context,
) {
	g := gin.Default()
	g.Use(middlewares.CORSMiddleware(), middlewares.JSONMiddleware(), RequestId())

	routes.InitHTTPRoute(g, mongodbClient, ctx)
	//useSSL, err := strconv.ParseBool(os.Getenv("USE_SSL"))
	addr := fmt.Sprintf(":%s", os.Getenv("MAIN_PORT"))

	//if err != nil || useSSL {
	//	g.RunTLS(addr, os.Getenv("PUBLIC_SSL_PATH"), os.Getenv("PRIVATE_SSL_PATH"))
	//} else {
	//	err = http.ListenAndServe(addr, g)
	//}

	//http.ListenAndServe(addr, g)
	g.Run(addr)
}

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Expose it for use in the application
		c.Set("RequestId", requestID)
		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
