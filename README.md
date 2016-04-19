# Golang Http服务测试范例

## 使用示例

``` go
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
```

## 获取范例

``` bash
$ git clone https://github.com/LyricTian/go-httptest-sample.git
$ go get -v github.com/smartystreets/goconvey/convey
```

## 运行范例

``` bash
$ cd $GOPATH/src/github.com/LyricTian/go-httptest-sample/tests
$ go test -v
```

## 输出

```
=== RUN   TestGetDefaultData
--- PASS: TestGetDefaultData (0.00s)
=== RUN   TestHelloController

  Test GetHello ✔✔


2 total assertions

--- PASS: TestHelloController (0.00s)
PASS
```