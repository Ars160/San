package main

import (
	"log"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
