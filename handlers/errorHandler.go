package handlers

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int){
w.WriteHeader(status)
if status == http.StatusNotFound {
	fmt.Fprint(w,"404 page not found")
}
}
