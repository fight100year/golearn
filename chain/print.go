package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello func")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		// fmt.Println("called - " + name)
		fmt.Fprintf(w, name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
