package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDD struct {
	client *mongo.Client
}

type IMongoDB interface {
	Client() *mongo.Client
}

type MongoDBParam struct {
	Host     string
	Port     int
	User     string
	Password string
	Local    bool
}

func NewMongoDB(param MongoDBParam) IMongoDB {
	var mongoURL string

	if param.Local {
		mongoURL = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	} else {
		mongoURL = fmt.Sprintf("mongodb+srv://%s:%s@%s", param.User, param.Password, param.Host)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))

	if err != nil {
		panic(err)
	}

	return &MongoDD{
		client: client,
	}
}

func (m *MongoDD) Client() *mongo.Client {
	return m.client
}
