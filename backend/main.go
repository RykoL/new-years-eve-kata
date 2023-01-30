package main

import (
	"github.com/RykoL/new-years-eve/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/wishes", handler.Router)

	http.ListenAndServe(":3000", mux)
}
