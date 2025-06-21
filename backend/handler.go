package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// HTTPハンドリング
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
	insertData := makeInsertData(data)
	err = insertDataToMongo(insertData)
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
