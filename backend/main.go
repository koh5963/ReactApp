package main

import (
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

func writeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ここ")
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
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":3001", nil)
}
