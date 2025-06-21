package main

// 受信するJSONデータの構造体
type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Message string
	Result  bool
}

type DbData struct {
	Id      int
	Message string
}
