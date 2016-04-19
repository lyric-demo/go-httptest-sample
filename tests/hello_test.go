package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/LyricTian/go-httptest-sample/router"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloController(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("Test GetHello", t, func() {
		r, _ := http.NewRequest(http.MethodGet, "/hello?name=Lyric", nil)
		http.DefaultServeMux.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, 200)
		So(w.Body.String(), ShouldEqual, "Hello,Lyric")
	})
}
