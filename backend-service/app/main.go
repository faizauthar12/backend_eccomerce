package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/faizauthar12/eccomerce/backend-service/app/handlers"
	"github.com/faizauthar12/eccomerce/global-utils/helper"
	"github.com/faizauthar12/eccomerce/global-utils/mongodb"
	"github.com/getsentry/sentry-go"
	envConfig "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	arg := os.Args[1]

	switch arg {
	case "main":
		mainWithoutArg()
		break
	default:
		mainWithoutArg()
		break
	}
}

func mainWithoutArg() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := envConfig.Load(".env"); err != nil {
		errStr := fmt.Sprintf(".env not load properly %s", err.Error())
		helper.SetSentryError(err, errStr, sentry.LevelError)
		panic(err)
	}

	ctx := context.Background()

	// mongoDB
	mongoDb := mongodb.NewMongoDB(mongodb.MongoDBParam{
		Local: true,
	})

	defer mongoDb.Client().Disconnect(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Printf("Starting Product Service HTTP Handler\n")
		handlers.MainHTTPHandler(mongoDb, ctx)
	}()

	wg.Wait()
}
