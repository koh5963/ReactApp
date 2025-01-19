package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// MongoDBに接続
func ConnectToMongo() error {
	mongoURI := "mongodb://admin:password@mongodb:27017" // docker-compose内のサービス名を指定
	clientOptions := options.Client().ApplyURI(mongoURI)

	// MongoDBに接続
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("Failed to connect to MongoDB: %v\n", err)
	}

	// 接続確認
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to ping MongoDB: %v\n", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
	return nil
}

// client の getter
func GetClient() *mongo.Client {
	return client
}
