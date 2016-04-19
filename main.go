package main

import (
	"log"
	"net/http"

	_ "github.com/LyricTian/go-httptest-sample/router"
)

func main() {
	log.Println("Server is running at :9090.")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
