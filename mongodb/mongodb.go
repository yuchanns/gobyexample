package mongodb

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

func (c *MongoClient) Close() error {
	return c.client.Disconnect(c.ctx)
}

func (c *MongoClient) GetConn() *mongo.Client {
	return c.client
}

func NewMongoClient() (*MongoClient, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI("mongodb://admin:admin@localhost:27017"))

	if err != nil {
		return nil, errors.Errorf("failed to connect mongodb: %v", err)
	}

	pingCtx, _ := context.WithTimeout(ctx, 5*time.Second)

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, errors.Errorf("failed to ping mongodb: %v", err)
	}

	return &MongoClient{
		client: client,
		ctx:    ctx,
	}, nil
}
