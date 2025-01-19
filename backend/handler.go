package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 受信するJSONデータの構造体
type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Message string
	Result  bool
}

// MongoDBにデータを格納する関数
func insertDataToMongo(data RequestData) error {
	client := GetClient() // dbパッケージからclientを取得

	if client == nil {
		return fmt.Errorf("Mongo client is not initialized")
	}

	db := client.Database("test_db")
	if db == nil {
		return fmt.Errorf("Failed to get MongoDB database")
	}

	collection := db.Collection("message")
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

func WriteHandler(w http.ResponseWriter, r *http.Request) {
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

	// ファイルに書き込む
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

	// レスポンス
	resData := ResponseData{Message: "Successfully wrote to file.", Result: true}
	jsonData, pErr := json.Marshal(resData)
	if pErr != nil {
		http.Error(w, pErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
