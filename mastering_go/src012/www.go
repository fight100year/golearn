package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
	fmt.Println("111")
	fmt.Fprintf(w, "111\n")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "the current time is: "
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
	fmt.Println("222")
	fmt.Fprintf(w, "222\n")
}

func main() {
	port := ":8080"

	if len(os.Args) != 1 {
		port = ":" + os.Args[1]
	}

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
