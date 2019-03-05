package main

import (
	"fmt"
	"net/http"
)

type h1 struct{}
type h2 struct{}

func (h *h1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello, %s", request.URL.Path[1:])
}

func (h *h2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello2, %s", request.URL.Path[1:])
}

func main() {
	hi1 := h1{}
	hi2 := h2{}

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/1/*", &hi1)
	http.Handle("/2/*", &hi2)
	server.ListenAndServe()
}
