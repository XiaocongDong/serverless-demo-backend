package main

import (
	"log"
	"net/http"

	analyze "github.com/XiaocongDong/serverless-demo-backend"
)

func main() {
	http.HandleFunc("/", analyze.Analyze)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
