package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
