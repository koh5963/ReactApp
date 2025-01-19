package main

import (
	"log"
	"net/http"
)

func main() {
	// MongoDBに接続
	err := ConnectToMongo()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return
	}

	// HTTPサーバー開始
	http.HandleFunc("/write", WriteHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
