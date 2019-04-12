package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func handler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Printf("server: %s\n", r.Host)
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
