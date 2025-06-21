package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// MongoDBにデータを格納する関数
func insertDataToMongo(data DbData) error {
	client := GetClient() // dbパッケージからclientを取得

	if client == nil {
		return fmt.Errorf("Mongo client is not initialized")
	}

	db := client.Database("test_db")
	if db == nil {
		return fmt.Errorf("Failed to get MongoDB database")
	}

	collection := db.Collection("request")
	if collection == nil {
		return fmt.Errorf("Mongo collection is nil")
	}

	// MongoDBにデータを挿入
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

// 登録用データ生成
func makeInsertData(data RequestData) DbData {
	fmt.Println("wow")
	var insertData DbData
	nextId := getNextId()
	insertData.Id = nextId
	insertData.Message = data.Message

	return insertData
}

// IDインクリメントを取得する関数
func getNextId() int {
	client := GetClient() // dbパッケージからclientを取得

	if client == nil {
		fmt.Println("Mongo client is not initialized")
		return 0
	}

	db := client.Database("test_db")
	if db == nil {
		fmt.Println("Failed to get MongoDB database")
		return 0
	}

	collection := db.Collection("request")
	if collection == nil {
		fmt.Println("Mongo collection is nil")
		return 0
	}

	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "maxValue", Value: bson.D{{Key: "$max", Value: "$id"}}},
		}}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Error running aggregation:", err)
		return 0
	}
	defer cursor.Close(ctx)

	// デフォルト値
	nextValue := 0

	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			fmt.Println("Error decoding result:", err)
			return 0
		}

		fmt.Printf("aggregation result: %+v\n", result)

		switch v := result["maxValue"].(type) {
		case int32:
			fmt.Println("int32")
			nextValue = int(v) + 1
		case int64:
			fmt.Println("int64")
			nextValue = int(v) + 1
		case float64:
			fmt.Println("other")
			nextValue = int(v) + 1
		default:
			fmt.Printf("unknown type for maxValue: %T\n", v)
		}
	}
	return nextValue
}
