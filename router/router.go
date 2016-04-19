package router

import (
	"net/http"

	"github.com/LyricTian/go-httptest-sample/controller"
)

func init() {
	http.HandleFunc("/", controller.GetDefaultData)
	http.HandleFunc("/hello", controller.GetHello)
}
