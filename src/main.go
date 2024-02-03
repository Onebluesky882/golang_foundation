package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {

	http.ListenAndServe(":8081", http.HandlerFunc(hello))
}
