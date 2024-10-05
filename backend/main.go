package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 受信するJSONデータの構造体
type RequestData struct {
	Message string `json:"message"`
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
	errFile := ioutil.WriteFile("./output.txt", []byte(data.Message), 0644)
	if errFile != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully wrote to file.")
}

func main() {
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":3001", nil)
}
