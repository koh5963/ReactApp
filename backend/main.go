package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func writeHandler(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	err := ioutil.WriteFile("C\\output.txt", []byte(message), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully wrote to file.")
}

func main() {
	fmt.Println("hello:P")
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":8080", nil)
}
