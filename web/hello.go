package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello, %s", request.URL.Path[1:])
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello2, %s", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/1/", handler2)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
