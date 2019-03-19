package main

import (
	"fmt"
	"net/http"
)

type h1 struct{}
type h2 struct{}

func (h *h1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, "hello, %s", request.URL.Path[1:])
	header := request.Header
	fmt.Fprintln(writer, header)
}

func (h *h2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(writer, "hello2, %s", request.URL.Path[1:])
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	hi1 := h1{}
	hi2 := h2{}

	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/abc", &hi1)
	http.Handle("/2", &hi2)
	server.ListenAndServe()
}
