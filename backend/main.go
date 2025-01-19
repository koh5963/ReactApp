package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 受信するJSONデータの構造体
type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Message string
	Result  bool
}

var client *mongo.Client

// MongoDBにデータを格納する関数
func insertDataToMongo(data RequestData) error {
	// clientが初期化されていなければ接続する
	if client == nil {
		// 環境変数から接続情報を取得
		mongoURI := "mongodb://admin:password@mongodb:27017" // docker-compose内のサービス名を指定
		clientOptions := options.Client().ApplyURI(mongoURI)

		// MongoDBに接続
		var err error
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return fmt.Errorf("Failed to connect to MongoDB: %v\n", err)
		}

		// 接続確認
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 時間を長く設定
		defer cancel()
		err = client.Ping(ctx, nil)
		if err != nil {
			return fmt.Errorf("Failed to ping MongoDB: %v\n", err)
		}

		fmt.Println("Successfully connected to MongoDB!")
	}

	// MongoDBデータベースの取得
	db := client.Database("test_db")
	if db == nil {
		return fmt.Errorf("Failed to get MongoDB database")
	}

	// コレクションの取得
	collection := db.Collection("message")
	if collection == nil {
		return fmt.Errorf("Mongo collection is nil!")
	}

	// MongoDBにデータを挿入
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORSを許可する
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := r.FormValue("message")
	fmt.Println(message)
	errFile := os.WriteFile("./output.txt", []byte(data.Message), 0644)
	if errFile != nil {
		http.Error(w, errFile.Error(), http.StatusInternalServerError)
		return
	}

	// MongoDBにデータを挿入
	err = insertDataToMongo(data)
	if err != nil {
		http.Error(w, "Failed to insert data into MongoDB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resData := ResponseData{Message: "Successfully wrote to file.", Result: true}
	jsonData, pErr := json.Marshal(resData)
	if pErr != nil {
		http.Error(w, pErr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, jsonData)
	w.Write(jsonData)
}

func main() {
	// 環境変数から接続情報を取得
	mongoURI := "mongodb://admin:password@mongodb:27017" // docker-compose内のサービス名を指定
	clientOptions := options.Client().ApplyURI(mongoURI)

	// MongoDBに接続
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		return // エラーが発生した場合は処理を中断
	}

	// 接続確認
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 時間を長く設定
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to ping MongoDB: %v\n", err)
		return // エラーが発生した場合は処理を中断
	}

	fmt.Println("Successfully connected to MongoDB!")
	// MongoDBデータベースの取得
	db := client.Database("test_db")
	if db == nil {
		fmt.Printf("Failed to get MongoDB database: %v\n", err)
		return
	}
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":3001", nil)
}
