package controller

import (
	"fmt"
	"net/http"
)

// GetHello Get hello test
func GetHello(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprint(w, "Hello,"+name)
}
