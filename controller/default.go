package controller

import (
	"fmt"
	"net/http"
)

// GetDefaultData Get default test
func GetDefaultData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
