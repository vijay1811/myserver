package main

import (
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world! from vijay waiting for ci")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, mux)
}
