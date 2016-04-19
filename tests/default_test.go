package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/LyricTian/go-httptest-sample/router"
)

func TestGetDefaultData(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Error("Request error")
		return
	}
	if v := w.Body.String(); v != "OK" {
		t.Error("Request data error:", v)
	}
}
